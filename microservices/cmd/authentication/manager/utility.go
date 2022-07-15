package manager

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"wartech-studio.com/monster-reacher/libraries/config"
	"wartech-studio.com/monster-reacher/libraries/database"
)

const ACCESS_TOKEN_LIFETIME = 5 * time.Minute
const ACCESS_TOKEN_EXTEND_AFTER = 1 * time.Minute
const ACCESS_TOKEN_EXTEND_MAX = 6

const NAME_DATABASE = "auth"
const NAME_TABLE = "token"

const HASH_DATA = "kmaADfas45asd16"

func getSecretKey(id string, ip string, platform string, timestamp int64) string {
	return fmt.Sprintf(`%s-%s-%s-%d`, id, ip, platform, timestamp)
}

func genAccessToken(secretKey string) string {
	hasher := sha256.New()
	hasher.Write([]byte(secretKey))
	return hex.EncodeToString(hasher.Sum([]byte(HASH_DATA)))
}

func getDriver() database.DBDriver {
	driver, err := database.NewMongoDBDriver(config.GetServiceConfig().Databases["mongodb"].Host, NAME_DATABASE, NAME_TABLE)

	if err != nil {
		panic(err)
	}

	return driver
}

func deleteToken(ctx context.Context, driver database.DBDriver, token string) error {
	if driver == nil {
		driver = getDriver()
		defer driver.Close()
	}

	return driver.DeleteOne(ctx, database.MongoDBSelectOneQueryFilterOne("access_token", token))
}
