package bff

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/protobuf/character"
	"wartech-studio.com/monster-reacher/libraries/protobuf/gateway"
	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
)

func LinkServiceToAccount(ctx context.Context, req *gateway.LinkServiceToAccountRequest) (empty *emptypb.Empty, err error) {
	userinfo, err := getUserInfoByServiceCode(ctx, req.GetServiceName(), req.GetServiceCode())
	if err != nil {
		return
	}
	owner, cc, err := requestService(ctx, req.GetToken(), config.GetNameConfig().MicroServiceName.Profile)
	if err != nil {
		return
	}
	defer cc.Close()
	c := profile.NewProfileClient(cc)
	_, err = c.AddServiceAuth(ctx, &profile.AddServiceAuthRequest{Id: owner, ServiceName: req.GetServiceName(), ServiceId: userinfo.GetId()})
	return
}

func GetProfileData(ctx context.Context, req *gateway.GetProfileDataRequest) (*gateway.GetProfileDataResponse, error) {
	owner, cc, err := requestService(ctx, req.GetToken(), config.GetNameConfig().MicroServiceName.Profile)
	if req.GetId() == "" {
		req.Id = owner
	}
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	c := profile.NewProfileClient(cc)
	resp, err := c.GetData(ctx, &profile.GetDataRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &gateway.GetProfileDataResponse{Data: resp.GetData()}, nil
}

func GetCharacterData(ctx context.Context, req *gateway.GetCharacterDataRequest) (*gateway.GetCharacterDataResponse, error) {
	owner, cc, err := requestService(ctx, req.GetToken(), config.GetNameConfig().MicroServiceName.Character)
	if req.GetId() == "" {
		req.Id = owner
	}
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	c := character.NewCharacterClient(cc)
	resp, err := c.GetData(ctx, &character.GetDataRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &gateway.GetCharacterDataResponse{Data: resp.GetData()}, nil
}

func SetCharacterName(ctx context.Context, req *gateway.SetCharacterNameRequest) (empty *emptypb.Empty, err error) {
	owner, cc, err := requestService(ctx, req.GetToken(), config.GetNameConfig().MicroServiceName.Character)
	if err != nil {
		return
	}
	defer cc.Close()

	c := character.NewCharacterClient(cc)
	_, err = c.SetName(ctx, &character.SetNameRequest{Id: owner, Name: req.GetName()})
	return
}

func SetCharacterMMR(ctx context.Context, req *gateway.SetCharacterMMRRequest) (empty *emptypb.Empty, err error) {
	owner, cc, err := requestService(ctx, req.GetToken(), config.GetNameConfig().MicroServiceName.Character)
	if err != nil {
		return
	}
	defer cc.Close()

	c := character.NewCharacterClient(cc)
	_, err = c.SetMMR(ctx, &character.SetMMRRequest{Id: owner, Mmr: req.GetMmr()})
	return
}

func IncrementCharacterEXP(ctx context.Context, req *gateway.IncrementCharacterEXPRequest) (empty *emptypb.Empty, err error) {
	owner, cc, err := requestService(ctx, req.GetToken(), config.GetNameConfig().MicroServiceName.Character)
	if err != nil {
		return
	}
	defer cc.Close()

	c := character.NewCharacterClient(cc)
	_, err = c.IncrementEXP(ctx, &character.IncrementEXPRequest{Id: owner, Exp: req.GetExp()})
	return
}
