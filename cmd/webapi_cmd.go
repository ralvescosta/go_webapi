package cmd

import (
	infra "webapi/pkg/infra"
)

func WebApi() error {
	configs := infra.WebApiConfig{
		AppHost: "0.0.0.0",
		AppPort: 3333,
	}

	return infra.Start(&configs)
}
