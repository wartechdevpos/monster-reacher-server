package manager

import (
	"context"
	"log"

	"github.com/google/uuid"
	"wartech-studio.com/monster-reacher/libraries/protobuf/services_discovery"
)

var services map[string]*services_discovery.ServiceInfo

func NewServerService() services_discovery.ServicesDiscoveryServer {
	service := services_discovery.NewServicesDiscoveryServer()
	service.RegisterHandler = Register
	service.HealthCheckHandler = HealthCheck
	service.CheckServiceIsOnlineHandler = CheckServiceIsOnline
	service.GetServicesHandler = GetServices
	return service
}

func Register(ctx context.Context, req *services_discovery.RegisterRequest) (*services_discovery.RegisterResponse, error) {
	for k := range services {
		if services[k].Host == req.GetHost() {
			if services[k].Name != req.GetService() {
				log.Printf("the service on host %s is switching from %s to %s\n", req.GetHost(), services[k].Name, req.GetService())
			}
			services[k].Name = req.GetService()
			log.Printf("the service %s on host %s was reconnect \n", req.GetService(), req.GetHost())
			return &services_discovery.RegisterResponse{Token: k}, nil
		}
	}
	uuid := uuid.New()
	services[uuid.String()] = &services_discovery.ServiceInfo{
		Name: req.GetService(),
		Host: req.GetHost(),
	}
	log.Printf("the service %s on host %s was register is success \n", req.GetService(), req.GetHost())
	return &services_discovery.RegisterResponse{Token: string(uuid.String())}, nil
}

func HealthCheck(stream services_discovery.ServicesDiscovery_HealthCheckServer) error {
	token := ""
	for {
		msg, err := stream.Recv()
		if _, found := services[msg.GetToken()]; !found {
			service := services[token]
			service.IsOnline = false
			log.Printf("the service %s on host %s is closed %s\n", service.Name, service.Host, err.Error())
			break
		}
		service := services[msg.GetToken()]
		if err != nil {
			log.Printf("the service %s on host %s is closed %s\n", service.Name, service.Host, err.Error())
			break
		}
		token = msg.GetToken()
		service.IsOnline = true
		stream.Send(&services_discovery.HealthCheckResponse{Success: true, Message: "ok"})
	}
	return nil
}

func CheckServiceIsOnline(ctx context.Context, req *services_discovery.CheckServiceIsOnlineRequest) (*services_discovery.CheckServiceIsOnlineResponse, error) {
	for k := range services {
		if services[k].GetName() == req.GetName() && services[k].GetIsOnline() {
			return &services_discovery.CheckServiceIsOnlineResponse{
				IsOnline: services[k].GetIsOnline(),
			}, nil
		}
	}
	return &services_discovery.CheckServiceIsOnlineResponse{IsOnline: false}, nil
}

func GetServices(context.Context, *services_discovery.GetServicesRequest) (*services_discovery.GetServicesresponse, error) {
	services := []*services_discovery.ServiceInfo{}

	for k := range services {
		services = append(services, services[k])
	}
	return &services_discovery.GetServicesresponse{
		Services: services,
	}, nil
}
