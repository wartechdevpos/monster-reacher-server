package bff

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
)

func GetUserData(ctx context.Context, token string, id string) (*data_schema.ProfileData, error) {
	ownerId, err := GetAuthTokenData(ctx, token)
	if err != nil {
		return nil, err
	}
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{
		config.GetNameConfig().MicroServiceName.Profile,
	})
	if !ok {
		err = errors.New("service profile is offline")
		return nil, err
	}

	if id == "" {
		id = ownerId
	}

	cc, err := grpc.Dial(serivces[config.GetNameConfig().MicroServiceName.Profile].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	c := profile.NewProfileClient(cc)
	resp, err := c.GetData(ctx, &profile.GetDataRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}
