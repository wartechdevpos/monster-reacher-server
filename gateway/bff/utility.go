package bff

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/authentication"
)

func errorServiceOffline(service string) error {
	return fmt.Errorf("service %s is offline", service)
}

func requestService(ctx context.Context, token string, service string) (owner string, conn *grpc.ClientConn, err error) {
	if owner, err = getAuthTokenOwner(ctx, token); err != nil {
		return
	}
	log.Println(owner, token, err)
	conn, err = grpcConn(service)
	return
}

func grpcConn(service string) (conn *grpc.ClientConn, err error) {
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{service})
	if !ok {
		err = errorServiceOffline(service)
		return
	}
	conn, err = grpc.Dial(serivces[service].Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return
}

func getAuthTokenOwner(ctx context.Context, token string) (id string, err error) {
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{
		config.GetNameConfig().MicroServiceName.Authentication,
	})
	if !ok {
		err = fmt.Errorf("service %s is offline", config.GetNameConfig().MicroServiceName.Authentication)
		return
	}

	cc, err := grpc.Dial(serivces[config.GetNameConfig().MicroServiceName.Authentication].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer cc.Close()

	c := authentication.NewAuthenticationClient(cc)
	resSignIn, err := c.SignIn(ctx, &authentication.SignInRequest{AccessToken: token})
	if err != nil {
		return
	}
	return resSignIn.GetId(), nil
}
