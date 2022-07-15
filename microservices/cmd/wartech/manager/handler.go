package manager

import (
	"context"

	"wartech-studio.com/monster-reacher/libraries/protobuf/wartech"
)

func NewServerService() wartech.WartechServer {
	service := wartech.NewWartechServer()
	service.RegisterHandler = Register
	service.AuthenticationHandler = Authentication
	service.CheckUserOrEmailHandler = CheckUserOrEmail
	// service.ForgottenPasswordHandler = ForgottenPassword
	// service.ChangePasswordHandler = ChangePassword
	service.GetUserHandler = GetUser
	return service
}

func Register(ctx context.Context, req *wartech.RegisterRequest) (*wartech.RegisterResponse, error) {
	id, err := internalRegister(ctx, req.GetUser(), req.GetEmail(), req.GetPassword(), req.GetBirthdayTimestamp())
	return &wartech.RegisterResponse{Id: id}, err
}

func Authentication(ctx context.Context, req *wartech.AuthenticationRequest) (*wartech.AuthenticationResponse, error) {
	data, err := internalAuthentication(ctx, req.GetUserOrEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &wartech.AuthenticationResponse{IsSuccess: data != nil, IsConfirmed: data.GetEmailConfirmed(), Id: data.GetId()}, err
}

func CheckUserOrEmail(ctx context.Context, req *wartech.CheckUserOrEmailRequest) (*wartech.CheckUserOrEmailResponse, error) {
	if _, found := internalCheckUser(ctx, req.GetUserOrEmail()); !found {
		if _, found := internalCheckEmail(ctx, req.GetUserOrEmail()); !found {
			return &wartech.CheckUserOrEmailResponse{IsValid: false}, nil
		}
	}
	return &wartech.CheckUserOrEmailResponse{IsValid: true}, nil
}

func GetUser(ctx context.Context, req *wartech.GetUserRequest) (*wartech.GetUserResponse, error) {
	data, err := internalGetUser(ctx, req.GetId())
	return &wartech.GetUserResponse{Data: data}, err
}
