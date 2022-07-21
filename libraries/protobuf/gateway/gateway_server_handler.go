package gateway

import context "context"

func (server *gatewayServer) SetAuthenticationHandler(handler func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error)) {
	server.AuthenticationHandler = handler
}

func (server *gatewayServer) SetWartechRegisterHandler(handler func(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterResponse, error)) {
	server.WartechRegisterHandler = handler
}

func (server *gatewayServer) SetGetProfileDataHandler(handler func(ctx context.Context, req *GetProfileDataRequest) (*GetProfileDataResponse, error)) {
	server.GetProfileDataHandler = handler
}
