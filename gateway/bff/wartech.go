package bff

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/gateway"
	"wartech-studio.com/monster-reacher/libraries/protobuf/wartech"
)

func WartechRegister(ctx context.Context, req *gateway.WartechRegisterRequest) (res *gateway.WartechRegisterResponse, err error) {
	res = &gateway.WartechRegisterResponse{}
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{
		config.GetNameConfig().MicroServiceName.Wartech,
	})
	if !ok {
		err = errors.New("service profile,authentication is offline")
		return
	}

	if err = checkWartechRegisterParam(req.GetUsername(), req.GetEmail(), req.GetPassword(), req.GetBirthday()); err != nil {
		return
	}

	cc, err := grpc.Dial(serivces[config.GetNameConfig().MicroServiceName.Wartech].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer cc.Close()

	c := wartech.NewWartechClient(cc)
	birthdayTime, _ := time.Parse("01-02-06", req.GetBirthday())
	resp, err := c.Register(ctx, &wartech.RegisterRequest{
		User:              req.GetUsername(),
		Email:             req.GetEmail(),
		Password:          req.GetPassword(),
		BirthdayTimestamp: timestamppb.New(birthdayTime),
	})
	res.IsSuccess = resp.GetIsSuccess()
	return
}

func checkWartechRegisterParam(username, email, password, birthday string) error {
	return nil
}
