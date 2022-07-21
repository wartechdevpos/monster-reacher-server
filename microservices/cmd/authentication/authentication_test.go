package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"wartech-studio.com/monster-reacher/libraries/protobuf/authentication"
	"wartech-studio.com/monster-reacher/microservices/cmd/authentication/manager"
)

func TestSignUp(t *testing.T) {
	service := manager.NewServerService()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := service.SignUp(ctx, &authentication.SignUpRequest{Id: "test"})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestSignIn(t *testing.T) {
	service := manager.NewServerService()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := service.SignIn(ctx, &authentication.SignInRequest{AccessToken: "6b6d614144666173343561736431363dea56288551e17157fbf8bdee2ccb2c72266bc2e84cb9d985994cf1da2b90ce"})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res.Id)
}
