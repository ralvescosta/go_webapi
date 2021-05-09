package cmd

import (
	infra "webapi/pkg/infra"
)

func WebApi() error {
	infra.SimpleServer()

	return nil
}
