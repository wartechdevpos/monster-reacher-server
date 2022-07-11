package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/gateway/services/gateway"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/healthcheck"
)

const SERVICES_NAME = "gateway"

var listenHost = fmt.Sprintf("%s:%d",
	config.WartechConfig().Services[SERVICES_NAME].Hosts[0],
	config.WartechConfig().Services[SERVICES_NAME].Ports[0])

func main() {
	initServicesDiscovery()
	server := grpc.NewServer()
	healthchecker := healthcheck.NewHealthCheckClient()
	go healthchecker.Start(SERVICES_NAME, listenHost)
	listener, err := net.Listen("tcp", listenHost)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	gateway.RegisterGatewayServer(server, gateway.NewGatewayServer())
	//reflection.Register(server)
	log.Println("gRPC server listening on " + listenHost)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func initServicesDiscovery() {
	servicesDiscoveryHost := fmt.Sprintf("%s:%d",
		config.WartechConfig().Services["services-discovery"].Hosts[0],
		config.WartechConfig().Services["services-discovery"].Ports[0])
	go api.ServicesDiscoveryCache.Start(servicesDiscoveryHost)
}
