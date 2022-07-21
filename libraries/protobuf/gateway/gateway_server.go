package gateway

import (
	context "context"
	"errors"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type gatewayServer struct {
	AuthenticationHandler        func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error)
	WartechRegisterHandler       func(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterResponse, error)
	GetProfileDataHandler        func(ctx context.Context, req *GetProfileDataRequest) (*GetProfileDataResponse, error)
	GetCharacterDataHandler      func(ctx context.Context, req *GetCharacterDataRequest) (*GetCharacterDataResponse, error)
	SetCharacterNameHandler      func(ctx context.Context, req *SetCharacterNameRequest) error
	SetCharacterMMRHandler       func(ctx context.Context, req *SetCharacterMMRRequest) error
	LinkServiceToAccountHandler  func(ctx context.Context, req *LinkServiceToAccountRequest) error
	IncrementCharacterEXPHandler func(ctx context.Context, req *IncrementCharacterEXPRequest) error
	AddToStorageHandler          func(ctx context.Context, req *AddToStorageRequest) error
}

//var checkGatewayServer GatewayServer = &gatewayServer{}

func NewGatewayServer() *gatewayServer {
	return &gatewayServer{}
}

func (server *gatewayServer) Authentication(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	if server.AuthenticationHandler == nil {
		return nil, errors.New("Authentication handler not implement")
	}
	return server.AuthenticationHandler(ctx, req)
}

func (server *gatewayServer) WartechRegister(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterResponse, error) {
	if server.WartechRegisterHandler == nil {
		return nil, errors.New("WartechRegister handler not implement")
	}
	return server.WartechRegisterHandler(ctx, req)
}

func (server *gatewayServer) GetProfileData(ctx context.Context, req *GetProfileDataRequest) (*GetProfileDataResponse, error) {
	if server.GetProfileDataHandler == nil {
		return nil, errors.New("GetProfileData handler not implement")
	}
	return server.GetProfileDataHandler(ctx, req)
}

func (server *gatewayServer) GetCharacterData(ctx context.Context, req *GetCharacterDataRequest) (*GetCharacterDataResponse, error) {
	if server.GetCharacterDataHandler == nil {
		return nil, errors.New("GetCharacterData handler not implement")
	}
	return server.GetCharacterDataHandler(ctx, req)
}
func (server *gatewayServer) SetCharacterName(ctx context.Context, req *SetCharacterNameRequest) (*emptypb.Empty, error) {
	if server.SetCharacterNameHandler == nil {
		return nil, errors.New("SetCharacterName handler not implement")
	}
	return &emptypb.Empty{}, server.SetCharacterNameHandler(ctx, req)
}
func (server *gatewayServer) SetCharacterMMR(ctx context.Context, req *SetCharacterMMRRequest) (*emptypb.Empty, error) {
	if server.SetCharacterMMRHandler == nil {
		return nil, errors.New("SetCharacterMMR handler not implement")
	}
	return &emptypb.Empty{}, server.SetCharacterMMRHandler(ctx, req)
}
func (server *gatewayServer) LinkServiceToAccount(ctx context.Context, req *LinkServiceToAccountRequest) (*emptypb.Empty, error) {
	if server.LinkServiceToAccountHandler == nil {
		return nil, errors.New("LinkServiceToAccount handler not implement")
	}
	return &emptypb.Empty{}, server.LinkServiceToAccountHandler(ctx, req)
}
func (server *gatewayServer) IncrementCharacterEXP(ctx context.Context, req *IncrementCharacterEXPRequest) (*emptypb.Empty, error) {
	if server.IncrementCharacterEXPHandler == nil {
		return nil, errors.New("IncrementCharacterEXP handler not implement")
	}
	return &emptypb.Empty{}, server.IncrementCharacterEXPHandler(ctx, req)
}
func (server *gatewayServer) AddToStorage(ctx context.Context, req *AddToStorageRequest) (*emptypb.Empty, error) {
	if server.AddToStorageHandler == nil {
		return nil, errors.New("AddToStorage handler not implement")
	}
	return &emptypb.Empty{}, server.AddToStorageHandler(ctx, req)
}

func (server *gatewayServer) mustEmbedUnimplementedGatewayServer() {}
