package profile

import (
	"context"
	"errors"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const NAME_DATABASE = "user"
const NAME_TABLE = "profile"

type profileServer struct {
	GetDataHandler           func(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error)
	AuthenticationHandler    func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error)
	RegisterHandler          func(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	AddServiceAuthHandler    func(ctx context.Context, req *AddServiceAuthRequest) (*emptypb.Empty, error)
	RemoveServiceAuthHandler func(ctx context.Context, req *RemoveServiceAuthRequest) (*emptypb.Empty, error)
}

func NewProfileServer() *profileServer {
	return &profileServer{}
}

func (server *profileServer) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	if server.GetDataHandler == nil {
		return nil, errors.New("GetData handler not implement")
	}
	return server.GetDataHandler(ctx, req)
}

func (server *profileServer) Authentication(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	if server.AuthenticationHandler == nil {
		return nil, errors.New("Authentication handler not implement")
	}
	return server.AuthenticationHandler(ctx, req)
}

func (server *profileServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	if server.RegisterHandler == nil {
		return nil, errors.New("Register handler not implement")
	}
	return server.RegisterHandler(ctx, req)
}

func (server *profileServer) AddServiceAuth(ctx context.Context, req *AddServiceAuthRequest) (*emptypb.Empty, error) {
	if server.AddServiceAuthHandler == nil {
		return nil, errors.New("AddServiceAuth handler not implement")
	}
	return server.AddServiceAuthHandler(ctx, req)
}

func (server *profileServer) RemoveServiceAuth(ctx context.Context, req *RemoveServiceAuthRequest) (*emptypb.Empty, error) {
	if server.RemoveServiceAuthHandler == nil {
		return nil, errors.New("RemoveServiceAuth handler not implement")
	}
	return server.RemoveServiceAuthHandler(ctx, req)
}

func (*profileServer) mustEmbedUnimplementedProfileServer() {}
