package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/healthcheck"
	"wartech-studio.com/monster-reacher/libraries/protobuf/authentication"
	"wartech-studio.com/monster-reacher/microservices/cmd/authentication/manager"
)

var SERVICES_NAME = config.GetNameConfig().MicroServiceName.Authentication

var listenHost = fmt.Sprintf("%s:%d",
	config.GetServiceConfig().Services[SERVICES_NAME].Hosts[0],
	config.GetServiceConfig().Services[SERVICES_NAME].Ports[0])

func main() {
	server := grpc.NewServer()
	healthchecker := healthcheck.NewHealthCheckClient()
	go healthchecker.Start(SERVICES_NAME, listenHost)
	listener, err := net.Listen("tcp", listenHost)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	authentication.RegisterAuthenticationServer(server, manager.NewServerService())
	//reflection.Register(server)
	log.Println("gRPC server listening on " + listenHost)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
