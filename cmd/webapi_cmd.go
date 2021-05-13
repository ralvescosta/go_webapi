package cmd

import (
	infra "webapi/pkg/infra"
)

func WebApi() error {
	configs := infra.WebApiConfig{
		DBHost:     "127.0.0.1",
		DBPort:     5432,
		DBUser:     "postgres",
		DBPassword: "123456",
		DBName:     "default",

		AppHost: "0.0.0.0",
		AppPort: 3333,
	}

	return infra.Start(&configs)
}
