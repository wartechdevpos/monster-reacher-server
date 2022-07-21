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
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
	"wartech-studio.com/monster-reacher/libraries/protobuf/gateway"
	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
	"wartech-studio.com/monster-reacher/libraries/protobuf/services_discovery"
)

func Authentication(ctx context.Context, req *gateway.AuthenticationRequest) (res *gateway.AuthenticationResponse, err error) {
	res = &gateway.AuthenticationResponse{}
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{
		config.GetNameConfig().MicroServiceName.Authentication,
		config.GetNameConfig().MicroServiceName.Profile})
	req.ServiceName = strings.ToLower(req.GetServiceName())
	if !ok {
		err = errors.New("service profile,authentication is offline")
		return
	}

	if req.GetServiceName() != "" {
		if res.Id, res.IsNew, err = register(ctx, serivces[config.GetNameConfig().MicroServiceName.Profile], req.GetServiceName(), req.GetServiceAuthCode()); err != nil {
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
	resSignUp, err := c.SignUp(ctx, &authentication.SignUpRequest{Id: res.Id})
	if err != nil {
		return
	}
	res.Token = resSignUp.GetAccessToken()
	return
}

func register(ctx context.Context, profileService *services_discovery.ServiceInfo, serviceName string, serviceAuthCode string) (id string, isNew bool, err error) {
	if profileService == nil {
		err = errors.New("services profile is offline")
		return
	}

	userinfo, err := getUserInfoByServiceCode(ctx, serviceName, serviceAuthCode)
	if err != nil {
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

	if result, err = c.Authentication(ctx, &profile.AuthenticationRequest{ServiceName: serviceName, ServiceId: userinfo.GetId()}); err == nil && result.GetId() != "" {
		id = result.GetId()
		return
	}
	var resultRegister *profile.RegisterResponse = nil
	if resultRegister, err = c.Register(ctx, &profile.RegisterRequest{
		ServiceName: serviceName,
		ServiceId:   userinfo.GetId(),
	}); err != nil {
		return
	}

	if resultRegister.GetId() == "" {
		err = fmt.Errorf("service %s id %s register fail", serviceName, userinfo.GetId())
		return
	}
	return resultRegister.GetId(), true, nil
}

func getUserInfoByServiceCode(ctx context.Context, serviceName string, serviceAuthCode string) (userInfo *data_schema.AuthenticationData, err error) {
	if serviceName == "" || serviceAuthCode == "" {
		err = errors.New("some a param is empty. please check params service_name,service_token")
		return
	}

	provider := oauth2.NewOAuth2Provider(serviceName)

	if provider == nil {
		err = errors.New("services " + serviceAuthCode + " not support")
		return
	}

	if err = provider.SubmitAuth(ctx, serviceAuthCode); err != nil {
		log.Println("SubmitAuth ", err.Error())
		err = errors.New("token is expired")
		return
	}

	if provider.GetData() == nil {
		err = errors.New("user info is empty")
		return
	}

	return provider.GetData(), nil
}
