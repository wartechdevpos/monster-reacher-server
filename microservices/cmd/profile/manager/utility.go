package manager

import (
	"context"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
)

const NAME_DATABASE = "user"
const NAME_TABLE = "profile"

func getDriver() database.DBDriver {
	driver, err := database.NewMongoDBDriver(config.GetServiceConfig().Databases["mongodb"].Host, NAME_DATABASE, NAME_TABLE)

	if err != nil {
		panic(err)
	}
	return driver
}
func getProfileData(ctx context.Context, driver database.DBDriver, filter interface{}) (*data_schema.ProfileData, error) {
	result := driver.SelectOne(ctx, filter)
	if err := database.MongoDBSelectOneResultGetError(result); err != nil {
		return nil, err
	}
	data := &data_schema.ProfileData{}
	if err := database.MongoDBDecodeResultToStruct(result, data); err != nil {
		return nil, nil
	}

	return data, nil
}
