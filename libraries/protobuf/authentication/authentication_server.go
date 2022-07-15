package authentication

import (
	context "context"
	"errors"
	"time"
)

const ACCESS_TOKEN_LIFETIME = 5 * time.Minute
const ACCESS_TOKEN_EXTEND_AFTER = 1 * time.Minute
const ACCESS_TOKEN_EXTEND_MAX = 6

const NAME_DATABASE = "auth"
const NAME_TABLE = "token"

const HASH_DATA = "kmaADfas45asd16"

type authenticationServer struct {
	SignUpHandler  func(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error)
	SignInHandler  func(ctx context.Context, req *SignInRequest) (*SignInResponse, error)
	SignOutHandler func(ctx context.Context, req *SignOutRequest) (*SignOutResponse, error)
}

func NewAuthenticationServer() *authenticationServer {
	return &authenticationServer{}
}

func (server *authenticationServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	if server.SignUpHandler == nil {
		return nil, errors.New("SignUp handler not implement")
	}
	return server.SignUpHandler(ctx, req)
}

func (server *authenticationServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	if server.SignInHandler == nil {
		return nil, errors.New("SignIn handler not implement")
	}
	return server.SignInHandler(ctx, req)
}

func (server *authenticationServer) SignOut(ctx context.Context, req *SignOutRequest) (*SignOutResponse, error) {
	if server.SignOutHandler == nil {
		return nil, errors.New("SignOut handler not implement")
	}
	return server.SignOutHandler(ctx, req)
}

func (*authenticationServer) mustEmbedUnimplementedAuthenticationServer() {}
