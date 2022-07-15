package manager

import (
	"context"
	"log"

	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
	"wartech-studio.com/monster-reacher/libraries/protobuf/profile"
)

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
	filter := database.MongoDBSelectOneQueryFilterOne("serviceAuth."+req.GetServiceName(), req.GetServiceId())
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

func UserIsValid(ctx context.Context, req *profile.UserIsValidRequest) (*profile.CheckProfileResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("auth.user", req.GetUser())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	return &profile.CheckProfileResponse{Id: data.GetId()}, nil
}
func NameIsValid(ctx context.Context, req *profile.NameIsValidRequest) (*profile.CheckProfileResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("name", req.GetName())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	return &profile.CheckProfileResponse{Id: data.GetId()}, nil
}
func ServiceIsValid(ctx context.Context, req *profile.ServiceIsValidRequest) (*profile.CheckProfileResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("serviceAuth."+req.GetName(), req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	return &profile.CheckProfileResponse{Id: data.GetId()}, nil
}
func ChangeName(ctx context.Context, req *profile.ChangeNameRequest) (*profile.SuccessResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	data.Name = req.GetNewName()
	err = driver.UpdateOne(ctx, filter, data)
	return &profile.SuccessResponse{Success: err == nil}, err
}

func AddServiceAuth(ctx context.Context, req *profile.AddServiceAuthRequest) (*profile.SuccessResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	data.Services[req.GetName()] = req.GetId()
	err = driver.UpdateOne(ctx, filter, data)
	return &profile.SuccessResponse{Success: err == nil}, err
}
func RemoveServiceAuth(ctx context.Context, req *profile.RemoveServiceAuthRequest) (*profile.SuccessResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	delete(data.Services, req.GetName())
	err = driver.UpdateOne(ctx, filter, data)
	return &profile.SuccessResponse{Success: err == nil}, err
}
func MergeData(ctx context.Context, req *profile.MergeDataRequest) (*profile.MergeDataResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetIdA())
	dataA, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	filter = database.MongoDBSelectOneQueryFilterOne("_id", req.GetIdB())
	dataB, err := getProfileData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}

	if len(dataA.GetServices()) == 0 {
		dataA.Services = dataB.Services
	}
	err = driver.UpdateOne(ctx, filter, dataA)
	if err != nil {
		return nil, err
	}
	err = driver.DeleteOne(ctx, filter)
	return &profile.MergeDataResponse{Id: dataA.GetId()}, err
}
