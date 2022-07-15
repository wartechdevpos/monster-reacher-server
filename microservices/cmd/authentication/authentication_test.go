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
