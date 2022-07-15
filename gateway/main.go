package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/gateway/bff"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/gateway"
)

const SERVICES_NAME = "gateway"

var listenHost = fmt.Sprintf("%s:%d",
	config.GetServiceConfig().Services[SERVICES_NAME].Hosts[0],
	config.GetServiceConfig().Services[SERVICES_NAME].Ports[0])

func main() {
	initServicesDiscovery()
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", listenHost)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	service := gateway.NewGatewayServer()

	service.AuthenticationHandler = Authentication
	service.WartechRegisterHandler = WartechRegister

	gateway.RegisterGatewayServer(server, service)
	reflection.Register(server)
	log.Println("gRPC server listening on " + listenHost)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func initServicesDiscovery() {
	serviceName := config.GetNameConfig().ServiceName.ServicesDiscovery
	servicesDiscoveryHost := fmt.Sprintf("%s:%d",
		config.GetServiceConfig().Services[serviceName].Hosts[0],
		config.GetServiceConfig().Services[serviceName].Ports[0])
	go api.ServicesDiscoveryCache.Start(servicesDiscoveryHost)
}

func Authentication(ctx context.Context, req *gateway.AuthenticationRequest) (*gateway.AuthenticationReasponse, error) {
	id, isNew, token, err := bff.Authentication(req.GetServiceName(), req.GetServiceAuthCode())
	if err != nil {
		return nil, err
	}
	return &gateway.AuthenticationReasponse{
		IsNew: isNew,
		Token: token,
		Id:    id,
	}, nil
}

func WartechRegister(ctx context.Context, req *gateway.WartechRegisterRequest) (*gateway.WartechRegisterReasponse, error) {
	success, err := bff.WartechRegister(req.GetUsername(), req.GetEmail(), req.GetPassword(), req.GetBirthday())
	return &gateway.WartechRegisterReasponse{
		IsSuccess: success,
	}, err
}
