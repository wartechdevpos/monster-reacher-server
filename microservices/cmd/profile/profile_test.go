package main

import (
	"context"
	"fmt"
	"testing"

	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
	"wartech-studio.com/monster-reacher/microservices/cmd/profile/manager"
)

func TestAddServiceAuth(t *testing.T) {
	service := manager.NewServerService()

	res, err := service.AddServiceAuth(context.Background(), &profile.AddServiceAuthRequest{
		Id:          "62d926cc284546b51820fb07",
		ServiceName: "facebook",
		ServiceId:   "001",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
