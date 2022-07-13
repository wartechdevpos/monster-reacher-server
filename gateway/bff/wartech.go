package bff

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"wartech-studio.com/monster-reacher/gateway/api"
	"wartech-studio.com/monster-reacher/gateway/services/wartech"
)

func WartechRegister(username, email, password, birthday string) (bool, error) {
	serivces, ok := api.ServicesDiscoveryCache.CheckRequireServices([]string{"wartech"})
	if !ok {
		return false, errors.New("service profile,authentication is offline")
	}

	if err := checkWartechRegisterParam(username, email, password, birthday); err != nil {
		return false, err
	}

	cc, err := grpc.Dial(serivces["wartech"].GetHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}
	defer cc.Close()

	c := wartech.NewWartechUserClient(cc)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	birthdayTime, _ := time.Parse("01-02-06", birthday)
	res, err := c.Register(ctx, &wartech.RegisterRequest{
		User:              username,
		Email:             email,
		Password:          password,
		BirthdayTimestamp: timestamppb.New(birthdayTime),
	})
	return res.GetIsSuccess(), err
}

func checkWartechRegisterParam(username, email, password, birthday string) error {
	return nil
}
