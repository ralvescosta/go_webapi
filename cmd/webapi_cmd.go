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

	os.Setenv("GO_ENV", viper.GetString("GO_ENV"))
	os.Setenv("APP_HOST", viper.GetString("APP_HOST"))
	os.Setenv("APP_PORT", viper.GetString("APP_PORT"))
	os.Setenv("GIN_MODE", viper.GetString("GIN_MODE"))

	os.Setenv("RSA_PRIVATE_KEY_DIR", viper.GetString("RSA_PRIVATE_KEY_DIR"))
	os.Setenv("RSA_PUBLIC_KEY_DIR", viper.GetString("RSA_PUBLIC_KEY_DIR"))
	os.Setenv("APP_ISSUER", viper.GetString("APP_ISSUER"))

	return infra.Start(&configs)
}
