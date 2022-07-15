package manager

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"
	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/database"
	"wartech-studio.com/monster-reacher/libraries/protobuf/data_schema"
)

const NAME_DATABASE = "wartech"
const NAME_TABLE = "accounts"

func getDriver() database.DBDriver {
	driver, err := database.NewMongoDBDriver(config.GetServiceConfig().Databases["mongodb"].Host, NAME_DATABASE, NAME_TABLE)

	if err != nil {
		panic(err)
	}
	return driver
}

func internalRegister(ctx context.Context, user string, email string, password string, birthday *timestamppb.Timestamp) (string, error) {
	driver := getDriver()
	defer driver.Close()
	data := &data_schema.WartechUserData{
		User:     user,
		Email:    email,
		Password: password,
		Birthday: birthday,
	}
	result, err := driver.PushOne(ctx, data)
	if err != nil {
		return "", err
	}
	return database.MongoDBDecodeResultToID(result), nil
}

func internalCheckUser(ctx context.Context, user string) (string, bool) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("user", user)
	data, err := getwartechUserData(ctx, driver, filter)
	if err != nil {
		return "", false
	}
	return data.GetId(), true
}

func internalCheckEmail(ctx context.Context, email string) (string, bool) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("email", email)
	data, err := getwartechUserData(ctx, driver, filter)
	if err != nil {
		return "", false
	}
	return data.GetId(), true
}

func internalGetUser(ctx context.Context, id string) (*data_schema.WartechUserData, error) {
	driver := getDriver()
	defer driver.Close()
	filter := database.MongoDBSelectOneQueryFilterOne("_id", id)
	data, err := getwartechUserData(ctx, driver, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func internalAuthentication(ctx context.Context, userOrEmail string, password string) (*data_schema.WartechUserData, error) {
	id := ""
	found := false

	if id, found = internalCheckUser(ctx, userOrEmail); !found {
		if id, found = internalCheckEmail(ctx, userOrEmail); !found {
			return nil, errors.New("not found")
		}
	}
	data, err := internalGetUser(ctx, id)

	if data.Password != password {
		return nil, errors.New("password do not match")
	}

	return data, err
}

func getwartechUserData(ctx context.Context, driver database.DBDriver, filter interface{}) (*data_schema.WartechUserData, error) {
	result := driver.SelectOne(ctx, filter)
	if err := database.MongoDBSelectOneResultGetError(result); err != nil {
		return nil, err
	}
	data := &data_schema.WartechUserData{}
	if err := database.MongoDBDecodeResultToStruct(result, data); err != nil {
		return nil, nil
	}

	return data, nil
}

func SelectizeData(data []byte) (*data_schema.WartechUserData, error) {
	w := &data_schema.WartechUserData{}
	err := json.Unmarshal(data, w)
	return w, err
}
