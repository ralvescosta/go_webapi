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

		AppHost: viper.GetString("APP_HOST"),
		AppPort: viper.GetInt("APP_PORT"),
	}

	return infra.Start(&configs)
}
