package gateway

import context "context"

func (server *gatewayServer) SetAuthenticationHandler(handler func(ctx context.Context, req *AuthenticationRequest) (*AuthenticationReasponse, error)) {
	server.AuthenticationHandler = handler
}

func (server *gatewayServer) SetWartechRegisterHandler(handler func(ctx context.Context, req *WartechRegisterRequest) (*WartechRegisterReasponse, error)) {
	server.WartechRegisterHandler = handler
}

func (server *gatewayServer) SetGetProfileDataHandler(handler func(ctx context.Context, req *GetProfileDataRequest) (*GetProfileDataReasponse, error)) {
	server.GetProfileDataHandler = handler
}
