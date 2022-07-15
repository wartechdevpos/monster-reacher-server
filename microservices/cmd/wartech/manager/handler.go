package manager

import (
	"context"

	"wartech-studio.com/monster-reacher/libraries/protobuf/wartech"
)

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
func ForgottenPassword(context.Context, *wartech.ForgottenPasswordRequest) (*wartech.ForgottenPasswordResponse, error) {
	return nil, nil
}
func ChangePassword(context.Context, *wartech.ChangePasswordRequest) (*wartech.ChangePasswordResponse, error) {
	return nil, nil
}
func GetUser(ctx context.Context, req *wartech.GetUserRequest) (*wartech.GetUserResponse, error) {
	data, err := internalGetUser(ctx, req.GetId())
	return &wartech.GetUserResponse{Data: data}, err
}
