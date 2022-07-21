package manager

import (
	"context"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
)

const NAME_DATABASE = "user"
const NAME_TABLE = "character"

func getDriver() database.DBDriver {
	driver, err := database.NewMongoDBDriver(config.GetServiceConfig().Databases["mongodb"].Host, NAME_DATABASE, NAME_TABLE)

	if err != nil {
		panic(err)
	}
	return driver
}
func getCharacterData(ctx context.Context, driver database.DBDriver, filter interface{}) (*data_schema.CharacterData, error) {
	result := driver.SelectOne(ctx, filter)
	if err := database.MongoDBSelectOneResultGetError(result); err != nil {
		return nil, err
	}
	data := &data_schema.CharacterData{}
	if err := database.MongoDBDecodeResultToStruct(result, data); err != nil {
		return nil, nil
	}

	return data, nil
}
