package wartech

import (
	context "context"
	"errors"
)

type wartechServer struct {
	RegisterHandler          func(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	AuthenticationHandler    func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error)
	CheckUserOrEmailHandler  func(ctx context.Context, req *CheckUserOrEmailRequest) (*CheckUserOrEmailResponse, error)
	ForgottenPasswordHandler func(context.Context, *ForgottenPasswordRequest) (*ForgottenPasswordResponse, error)
	ChangePasswordHandler    func(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	GetUserHandler           func(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error)
	ConfirmEmailHandler      func(context.Context, *ConfirmEmailRequest) (*ConfirmEmailResponse, error)
}

func NewWartechServer() *wartechServer {
	return &wartechServer{}
}

func (server *wartechServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	if server.RegisterHandler == nil {
		return nil, errors.New("Register handler not implement")
	}
	return server.RegisterHandler(ctx, req)
}
func (server *wartechServer) Authentication(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	if server.AuthenticationHandler == nil {
		return nil, errors.New("Authentication handler not implement")
	}
	return server.AuthenticationHandler(ctx, req)
}
func (server *wartechServer) CheckUserOrEmail(ctx context.Context, req *CheckUserOrEmailRequest) (*CheckUserOrEmailResponse, error) {
	if server.CheckUserOrEmailHandler == nil {
		return nil, errors.New("CheckUserOrEmail handler not implement")
	}
	return server.CheckUserOrEmailHandler(ctx, req)
}
func (server *wartechServer) ForgottenPassword(ctx context.Context, req *ForgottenPasswordRequest) (*ForgottenPasswordResponse, error) {
	if server.ForgottenPasswordHandler == nil {
		return nil, errors.New("ForgottenPassword handler not implement")
	}
	return server.ForgottenPasswordHandler(ctx, req)
}
func (server *wartechServer) ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	if server.ChangePasswordHandler == nil {
		return nil, errors.New("ChangePassword handler not implement")
	}
	return server.ChangePasswordHandler(ctx, req)
}
func (server *wartechServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	if server.GetUserHandler == nil {
		return nil, errors.New("GetUser handler not implement")
	}
	return server.GetUserHandler(ctx, req)
}
func (server *wartechServer) ConfirmEmail(ctx context.Context, req *ConfirmEmailRequest) (*ConfirmEmailResponse, error) {
	if server.ConfirmEmailHandler == nil {
		return nil, errors.New("ConfirmEmail handler not implement")
	}
	return server.ConfirmEmailHandler(ctx, req)
}
func (*wartechServer) mustEmbedUnimplementedWartechServer() {}
