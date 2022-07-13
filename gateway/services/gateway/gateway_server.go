package gateway

import (
	context "context"

	"wartech-studio.com/monster-reacher/gateway/bff"
)

type gatewayServer struct{}

func NewGatewayServer() GatewayServer {
	return &gatewayServer{}
}

func (*gatewayServer) Authentication(ctx context.Context, req *AuthenticationRequest) (*AuthenticationReasponse, error) {
	id, isNew, token, err := bff.Authentication(req.GetUser(), req.GetPassword(), req.GetEmail(), req.GetServiceName(), req.GetServiceAuthCode())
	if err != nil {
		return nil, err
	}
	return &AuthenticationReasponse{
		IsNew: isNew,
		Token: token,
		Id:    id,
	}, nil
}

func (*gatewayServer) WartechRegister(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterReasponse, error) {
	success, err := bff.WartechRegister(req.GetUsername(), req.GetEmail(), req.GetPassword(), req.GetBirthday())
	return &WartechRegisterReasponse{
		IsSuccess: success,
	}, err
}

func (*gatewayServer) mustEmbedUnimplementedGatewayServer() {}
