package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/services_discovery"
	"wartech-studio.com/monster-reacher/services-discovery/manager"
)

var SERVICES_NAME = config.GetNameConfig().ServiceName.ServiceDiscovery

var listenHost = fmt.Sprintf("%s:%d",
	config.GetServiceConfig().Services[SERVICES_NAME].Hosts[0],
	config.GetServiceConfig().Services[SERVICES_NAME].Ports[0])

func main() {
	server := grpc.NewServer()

	listener, err := net.Listen("tcp", listenHost)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	services_discovery.RegisterServicesDiscoveryServer(server, manager.NewServerService())

	log.Println("services-discovery listening on " + listenHost)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
