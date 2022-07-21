package manager

import (
	"context"

	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/character"
)

func NewServerService() character.CharacterServer {
	service := character.NewCharacterServer()
	service.GetDataHandler = GetData
	service.SetNameHandler = SetName
	service.SetMMRHandler = SetMMR
	service.IncrementEXPHandler = IncrementEXP
	return service
}

func GetData(ctx context.Context, req *character.GetDataRequest) (*character.GetDataResponse, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	data, err := getCharacterData(ctx, driver, filter)
	return &character.GetDataResponse{Data: data}, err
}
func SetName(ctx context.Context, req *character.SetNameRequest) error {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	return driver.UpdateSpecific(ctx, filter, "name", req.GetName())
}
func SetMMR(ctx context.Context, req *character.SetMMRRequest) error {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	return driver.UpdateSpecific(ctx, filter, "mmr", req.GetMmr())
}
func IncrementEXP(ctx context.Context, req *character.IncrementEXPRequest) error {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", req.GetId())
	return driver.IncrementValue(ctx, filter, "exp", req.GetExp())
}
