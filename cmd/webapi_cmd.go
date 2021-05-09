package cmd

import (
	infra "webapi/pkg/infra"
)

func WebApi() error {
	configs := infra.WebApiConfig{}

	return infra.Server(&configs)
}
