package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/gateway/manager"
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

	gateway.RegisterGatewayServer(server, manager.NewServerService())
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
