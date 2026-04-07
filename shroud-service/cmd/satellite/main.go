package main

import (
	"services/internal/api"
	"services/internal/config"
	"services/internal/database/redisdb"
)

func main() {
	environment, err := config.Load[config.SatelliteEnvironment]()
	if err != nil {
		panic(err)
	}

	redisdb.LoadConfig(environment.RedisUri, environment.RedisPassword)

	api.LoadAuthSecret(environment.UserAuthSecret)
	api.LoadUploadConf(environment.UserUploadSecret, environment.UserUploadUri)

	api.StartRouter()
}
