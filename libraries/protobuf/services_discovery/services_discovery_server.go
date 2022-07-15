package services_discovery

import (
	context "context"
	"errors"
)

type servicesDiscoveryServer struct {
	RegisterHandler             func(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	HealthCheckHandler          func(stream ServicesDiscovery_HealthCheckServer) error
	CheckServiceIsOnlineHandler func(ctx context.Context, req *CheckServiceIsOnlineRequest) (*CheckServiceIsOnlineResponse, error)
	GetServicesHandler          func(ctx context.Context, req *GetServicesRequest) (*GetServicesresponse, error)
}

func NewServicesDiscoveryServer() *servicesDiscoveryServer {
	return &servicesDiscoveryServer{}
}

func (server *servicesDiscoveryServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	if server.RegisterHandler == nil {
		return nil, errors.New("Register handler not implement")
	}
	return server.RegisterHandler(ctx, req)
}

func (server *servicesDiscoveryServer) HealthCheck(stream ServicesDiscovery_HealthCheckServer) error {
	if server.HealthCheckHandler == nil {
		return errors.New("HealthCheck handler not implement")
	}
	return server.HealthCheckHandler(stream)
}
func (server *servicesDiscoveryServer) CheckServiceIsOnline(ctx context.Context, req *CheckServiceIsOnlineRequest) (*CheckServiceIsOnlineResponse, error) {
	if server.CheckServiceIsOnlineHandler == nil {
		return nil, errors.New("CheckServiceIsOnline handler not implement")
	}
	return server.CheckServiceIsOnlineHandler(ctx, req)
}
func (server *servicesDiscoveryServer) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesresponse, error) {
	if server.GetServicesHandler == nil {
		return nil, errors.New("GetServices handler not implement")
	}
	return server.GetServicesHandler(ctx, req)
}
func (*servicesDiscoveryServer) mustEmbedUnimplementedServicesDiscoveryServer() {}
