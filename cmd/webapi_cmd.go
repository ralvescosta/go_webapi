package cmd

import (
	infra "webapi/pkg/infra"
	"webapi/pkg/infra/env"

	"github.com/spf13/viper"
)

func WebApi() error {
	env.ConfigEnvs()

	configs := infra.WebApiConfig{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),

		Env:     viper.GetString("GO_ENV"),
		AppHost: viper.GetString("APP_HOST"),
		AppPort: viper.GetInt("APP_PORT"),
		GinMode: viper.GetString("GIN_MODE"),

		WebApiReqsLog: viper.GetString("WEBAPI_REQS_LOG"),
	}

	return infra.Start(&configs)
}
