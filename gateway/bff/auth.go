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
	"wartech-studio.com/monster-reacher/gateway/services/authentication"
	"wartech-studio.com/monster-reacher/gateway/services/profile"
	"wartech-studio.com/monster-reacher/gateway/services/services_discovery"
)

func Authentication(user string, password string, email string, serviceName string, serviceAuthCode string) (id string, isNew bool, token string, err error) {
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{"authentication", "profile"})
	serviceName = strings.ToLower(serviceName)
	if !ok {
		err = errors.New("service profile,authentication is offline")
		return
	}

	if user != "" {
		if id, isNew, err = RegisterByUser(serivces["profile"], user, password, email); err != nil {
			return
		}
	} else if serviceName != "" {
		if id, isNew, err = RegisterByService(serivces["profile"], serviceName, serviceAuthCode); err != nil {
			return
		}
	} else {
		err = errors.New("please check request parameter")
		return
	}

	cc, err := grpc.Dial(serivces["authentication"].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func RegisterByUser(profileService *services_discovery.ServiceInfo, user string, password string, email string) (id string, isNew bool, err error) {

	if user == "" || password == "" || email == "" {
		err = fmt.Errorf("some a param is empty. please check params user,password,email")
		return
	}

	if profileService == nil {
		err = errors.New("services profile is offline")
		return
	}

	cc, err := grpc.Dial(profileService.GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err = fmt.Errorf("serivces Dial is error %s", err.Error())
		return
	}
	defer cc.Close()

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	c := profile.NewProfileClient(cc)
	var result *profile.CheckProfileResponse
	if result, err = c.Authentication(ctx, &profile.AuthenticationRequest{User: user, Password: password}); err == nil && result.GetId() != "" {
		id = result.GetId()
		return
	}

	var resultRegister *profile.RegisterResponse = nil
	if resultRegister, err = c.Register(ctx, &profile.RegisterRequest{
		User:     user,
		Password: password,
		Email:    email,
	}); err != nil {
		return
	}
	return resultRegister.GetId(), true, nil
}

func RegisterByService(profileService *services_discovery.ServiceInfo, serviceName string, serviceAuthCode string) (id string, isNew bool, err error) {

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

	var result *profile.CheckProfileResponse = nil
	if result, err = c.AuthenticationByService(ctx, &profile.AuthenticationByServiceRequest{Name: serviceName, Id: provider.GetData().ID}); err == nil && result.GetId() != "" {
		id = result.GetId()
		return
	}
	var resultRegister *profile.RegisterResponse = nil
	if resultRegister, err = c.RegisterByService(ctx, &profile.RegisterByServiceRequest{
		Name: provider.GetServiceName(),
		Id:   provider.GetData().ID,
	}); err != nil {
		return
	}

	if resultRegister.GetId() == "" {
		err = fmt.Errorf("service %s id %s register fail", provider.GetServiceName(), provider.GetData().ID)
		return
	}
	return resultRegister.GetId(), true, nil
}
