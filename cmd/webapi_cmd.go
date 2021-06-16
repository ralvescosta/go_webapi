package cmd

import (
	"os"
	infra "webapi/pkg/infra"
	"webapi/pkg/infra/env"

	"github.com/spf13/viper"
)

func WebApi() error {
	env.ConfigEnvs()

	configs := infra.WebApiConfig{
		SQLDbHost:     viper.GetString("SQL_DB_HOST"),
		SQLDbPort:     viper.GetInt("SQL_DB_PORT"),
		SQLDbUser:     viper.GetString("SQL_DB_USER"),
		SQLDbPassword: viper.GetString("SQL_DB_PASSWORD"),
		SQLDbName:     viper.GetString("SQL_DB_NAME"),

		MongoDbHost:     viper.GetString("MONGO_DB_HOST"),
		MongoDbPort:     viper.GetInt("MONGO_DB_PORT"),
		MongoDbUser:     viper.GetString("MONGO_DB_USER"),
		MongoDbPassword: viper.GetString("MOND_DB_PASSWORD"),
		MongoDbName:     viper.GetString("MONGO_DB_NAME"),

		Env:     viper.GetString("GO_ENV"),
		AppHost: viper.GetString("APP_HOST"),
		AppPort: viper.GetInt("APP_PORT"),
		GinMode: viper.GetString("GIN_MODE"),

		WebApiReqsLog: viper.GetString("WEBAPI_REQS_LOG"),
	}

	os.Setenv("GO_ENV", viper.GetString("GO_ENV"))
	os.Setenv("APP_HOST", viper.GetString("APP_HOST"))
	os.Setenv("APP_PORT", viper.GetString("APP_PORT"))
	os.Setenv("GIN_MODE", viper.GetString("GIN_MODE"))

	os.Setenv("RSA_PRIVATE_KEY_DIR", viper.GetString("RSA_PRIVATE_KEY_DIR"))
	os.Setenv("RSA_PUBLIC_KEY_DIR", viper.GetString("RSA_PUBLIC_KEY_DIR"))
	os.Setenv("APP_ISSUER", viper.GetString("APP_ISSUER"))

	return infra.Start(&configs)
}
