package profile

import (
	"context"
	"errors"
)

const NAME_DATABASE = "user"
const NAME_TABLE = "profile"

type profileServer struct {
	GetDataHandler           func(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error)
	AuthenticationHandler    func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error)
	RegisterHandler          func(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	UserIsValidHandler       func(ctx context.Context, req *UserIsValidRequest) (*CheckProfileResponse, error)
	NameIsValidHandler       func(ctx context.Context, req *NameIsValidRequest) (*CheckProfileResponse, error)
	ServiceIsValidHandler    func(ctx context.Context, req *ServiceIsValidRequest) (*CheckProfileResponse, error)
	ChangeNameHandler        func(ctx context.Context, req *ChangeNameRequest) (*SuccessResponse, error)
	AddServiceAuthHandler    func(ctx context.Context, req *AddServiceAuthRequest) (*SuccessResponse, error)
	RemoveServiceAuthHandler func(ctx context.Context, req *RemoveServiceAuthRequest) (*SuccessResponse, error)
	MergeDataHandler         func(ctx context.Context, req *MergeDataRequest) (*MergeDataResponse, error)
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

func (server *profileServer) UserIsValid(ctx context.Context, req *UserIsValidRequest) (*CheckProfileResponse, error) {
	if server.UserIsValidHandler == nil {
		return nil, errors.New("UserIsValid handler not implement")
	}
	return server.UserIsValidHandler(ctx, req)
}
func (server *profileServer) NameIsValid(ctx context.Context, req *NameIsValidRequest) (*CheckProfileResponse, error) {
	if server.NameIsValidHandler == nil {
		return nil, errors.New("NameIsValid handler not implement")
	}
	return server.NameIsValidHandler(ctx, req)
}
func (server *profileServer) ServiceIsValid(ctx context.Context, req *ServiceIsValidRequest) (*CheckProfileResponse, error) {
	if server.ServiceIsValidHandler == nil {
		return nil, errors.New("ServiceIsValid handler not implement")
	}
	return server.ServiceIsValidHandler(ctx, req)
}
func (server *profileServer) ChangeName(ctx context.Context, req *ChangeNameRequest) (*SuccessResponse, error) {
	if server.ChangeNameHandler == nil {
		return nil, errors.New("ChangeName handler not implement")
	}
	return server.ChangeNameHandler(ctx, req)
}

func (server *profileServer) AddServiceAuth(ctx context.Context, req *AddServiceAuthRequest) (*SuccessResponse, error) {
	if server.AddServiceAuthHandler == nil {
		return nil, errors.New("AddServiceAuth handler not implement")
	}
	return server.AddServiceAuthHandler(ctx, req)
}
func (server *profileServer) RemoveServiceAuth(ctx context.Context, req *RemoveServiceAuthRequest) (*SuccessResponse, error) {
	if server.RemoveServiceAuthHandler == nil {
		return nil, errors.New("RemoveServiceAuth handler not implement")
	}
	return server.RemoveServiceAuthHandler(ctx, req)
}
func (server *profileServer) MergeData(ctx context.Context, req *MergeDataRequest) (*MergeDataResponse, error) {
	if server.MergeDataHandler == nil {
		return nil, errors.New("MergeData handler not implement")
	}
	return server.MergeDataHandler(ctx, req)
}
func (*profileServer) mustEmbedUnimplementedProfileServer() {}
