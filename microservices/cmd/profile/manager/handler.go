package manager

import (
	"context"
	"log"

	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
)

func NewServerService() profile.ProfileServer {
	service := profile.NewProfileServer()
	service.GetDataHandler = GetData
	service.AuthenticationHandler = Authentication
	service.RegisterHandler = Register
	service.AddServiceAuthHandler = AddServiceAuth
	service.RemoveServiceAuthHandler = RemoveServiceAuth
	return service
}

func GetData(ctx context.Context, req *profile.GetDataRequest) (*profile.GetDataResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	return &profile.GetDataResponse{Data: data}, err
}

func Authentication(ctx context.Context, req *profile.AuthenticationRequest) (*profile.AuthenticationResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("services."+req.GetServiceName(), req.GetServiceId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	return &profile.AuthenticationResponse{Id: data.GetId()}, nil
}

func Register(ctx context.Context, req *profile.RegisterRequest) (*profile.RegisterResponse, error) {
	driver := getDriver()
	defer driver.Close()
	services := make(map[string]string)
	services[req.GetServiceName()] = req.GetServiceId()
	data := &data_schema.ProfileData{
		Services: services,
	}
	result, err := driver.PushOne(ctx, data)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	return &profile.RegisterResponse{Id: database.MongoDBDecodeResultToID(result)}, nil
}

func AddServiceAuth(ctx context.Context, req *profile.AddServiceAuthRequest) error {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return err
	}
	data.Services[req.GetServiceName()] = req.GetServiceId()
	return driver.UpdateOne(ctx, filter, data)
}

func RemoveServiceAuth(ctx context.Context, req *profile.RemoveServiceAuthRequest) error {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return err
	}
	delete(data.Services, req.GetServiceName())
	return driver.UpdateOne(ctx, filter, data)
}
