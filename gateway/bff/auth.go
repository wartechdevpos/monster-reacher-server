package bff

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/gateway/api/oauth2"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/authentication"
	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
	"wartech-studio.com/monster-reacher/libraries/protobuf/services_discovery"
)

func Authentication(serviceName string, serviceAuthCode string) (id string, isNew bool, token string, err error) {
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{
		config.GetNameConfig().MicroServiceName.Authentication,
		config.GetNameConfig().MicroServiceName.Profile})
	serviceName = strings.ToLower(serviceName)
	if !ok {
		err = errors.New("service profile,authentication is offline")
		return
	}

	if serviceName != "" {
		if id, isNew, err = Register(serivces[config.GetNameConfig().MicroServiceName.Profile], serviceName, serviceAuthCode); err != nil {
			return
		}
	} else {
		err = errors.New("please check request parameter")
		return
	}

	cc, err := grpc.Dial(serivces[config.GetNameConfig().MicroServiceName.Authentication].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer cc.Close()

	c := authentication.NewAuthenticationClient(cc)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	resSignUp, err := c.SignUp(ctx, &authentication.SignUpRequest{Id: id})
	if err != nil {
		return
	}
	token = resSignUp.GetAccessToken()
	return
}

func Register(profileService *services_discovery.ServiceInfo, serviceName string, serviceAuthCode string) (id string, isNew bool, err error) {

	if serviceName == "" || serviceAuthCode == "" {
		err = errors.New("some a param is empty. please check params service_name,service_token")
		return
	}

	provider := oauth2.NewOAuth2Provider(serviceName)

	if provider == nil {
		err = errors.New("services " + serviceAuthCode + " not support")
		return
	}

	if profileService == nil {
		err = errors.New("services profile is offline")
		return
	}

	if err = provider.SubmitAuth(serviceAuthCode); err != nil {
		log.Println("SubmitAuth ", err.Error())
		err = errors.New("token is expired")
		return
	}

	if provider.GetData() == nil {
		err = errors.New("user info is empty")
		return
	}

	cc, err := grpc.Dial(profileService.GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer cc.Close()

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	c := profile.NewProfileClient(cc)

	var result *profile.AuthenticationResponse = nil

	if result, err = c.Authentication(ctx, &profile.AuthenticationRequest{ServiceName: serviceName, ServiceId: provider.GetData().GetId()}); err == nil && result.GetId() != "" {
		id = result.GetId()
		return
	}
	var resultRegister *profile.RegisterResponse = nil
	if resultRegister, err = c.Register(ctx, &profile.RegisterRequest{
		ServiceName: provider.GetServiceName(),
		ServiceId:   provider.GetData().GetId(),
	}); err != nil {
		return
	}

	if resultRegister.GetId() == "" {
		err = fmt.Errorf("service %s id %s register fail", provider.GetServiceName(), provider.GetData().GetId())
		return
	}
	return resultRegister.GetId(), true, nil
}

func GetAuthTokenData(ctx context.Context, token string) (id string, err error) {
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{
		config.GetNameConfig().MicroServiceName.Authentication,
	})
	if !ok {
		err = errors.New("service authentication is offline")
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
