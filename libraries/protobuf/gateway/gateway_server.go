package gateway

import (
	context "context"
	"errors"
)

type gatewayServer struct {
	AuthenticationHandler  func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationReasponse, error)
	WartechRegisterHandler func(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterReasponse, error)
	GetProfileDataHandler  func(ctx context.Context, req *GetProfileDataRequest) (*GetProfileDataReasponse, error)
}

func NewGatewayServer() *gatewayServer {
	return &gatewayServer{}
}

func (server *gatewayServer) Authentication(ctx context.Context, req *AuthenticationRequest) (*AuthenticationReasponse, error) {
	if server.AuthenticationHandler == nil {
		return nil, errors.New("Authentication handler not implement")
	}
	return server.AuthenticationHandler(ctx, req)
}

func (server *gatewayServer) WartechRegister(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterReasponse, error) {
	if server.WartechRegisterHandler == nil {
		return nil, errors.New("WartechRegister handler not implement")
	}
	return server.WartechRegisterHandler(ctx, req)
}

func (server *gatewayServer) GetProfileData(ctx context.Context, req *GetProfileDataRequest) (*GetProfileDataReasponse, error) {
	if server.GetProfileDataHandler == nil {
		return nil, errors.New("GetProfileData handler not implement")
	}
	return server.GetProfileDataHandler(ctx, req)
}

func (server *gatewayServer) mustEmbedUnimplementedGatewayServer() {}
