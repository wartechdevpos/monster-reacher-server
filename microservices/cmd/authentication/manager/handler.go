package manager

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/authentication"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
)

func SignUp(ctx context.Context, req *authentication.SignUpRequest) (*authentication.SignUpResponse, error) {
	dbDriver := getDriver()
	defer dbDriver.Close()

	dbDriver.DeleteOne(ctx, database.MongoDBSelectOneQueryFilterOne("id", req.GetId()))

	createTimestamp := time.Now().UTC().Unix()
	secretKey := getSecretKey(req.Id, req.Ip, req.Platform, createTimestamp)
	accessToken := genAccessToken(secretKey)

	_, err := dbDriver.PushOne(ctx, data_schema.AuthenticationData{
		Id:              req.Id,
		CreateTimestamp: timestamppb.New(time.Now().UTC()),
		AccessToken:     accessToken,
		ExtendCount:     0,
	})

	if err != nil {
		return nil, err
	}

	log.Printf("generate new access token from id %s token %s date %s \n", req.Id, accessToken, time.Unix(createTimestamp, 0).Format(time.UnixDate))

	return &authentication.SignUpResponse{
		AccessToken: accessToken,
	}, nil
}

func SignIn(ctx context.Context, req *authentication.SignInRequest) (*authentication.SignInResponse, error) {
	dbDriver := getDriver()
	defer dbDriver.Close()

	result := dbDriver.SelectOne(ctx, database.MongoDBSelectOneQueryFilterOne("access_token", req.AccessToken))

	if err := database.MongoDBSelectOneResultGetError(result); err != nil {
		return &authentication.SignInResponse{IsValid: false}, nil
	}

	authData := &data_schema.AuthenticationData{}

	if err := database.MongoDBDecodeResultToStruct(result, authData); err != nil {
		return &authentication.SignInResponse{IsValid: false}, err
	}

	timeout := authData.CreateTimestamp.GetSeconds() + int64(ACCESS_TOKEN_LIFETIME.Seconds()*float64(authData.ExtendCount))

	if time.Now().UTC().Unix()-timeout <= 0 {
		return &authentication.SignInResponse{IsValid: true}, nil
	}

	if time.Now().UTC().Unix()-timeout > int64(ACCESS_TOKEN_LIFETIME.Seconds()) {
		deleteToken(ctx, dbDriver, req.AccessToken)
		return &authentication.SignInResponse{IsValid: false}, nil
	}

	if authData.ExtendCount > ACCESS_TOKEN_EXTEND_MAX {
		deleteToken(ctx, dbDriver, req.AccessToken)
		return &authentication.SignInResponse{IsValid: false}, nil
	}

	authData.ExtendCount = authData.ExtendCount + 1

	if err := dbDriver.UpdateOne(ctx, database.MongoDBSelectOneQueryFilterOne("access_token", req.AccessToken), authData); err != nil {
		return &authentication.SignInResponse{IsValid: true}, err
	}

	return &authentication.SignInResponse{IsValid: true, Id: authData.GetId()}, nil
}

func SignOut(ctx context.Context, req *authentication.SignOutRequest) (*authentication.SignOutResponse, error) {
	return &authentication.SignOutResponse{}, deleteToken(ctx, nil, req.AccessToken)
}
