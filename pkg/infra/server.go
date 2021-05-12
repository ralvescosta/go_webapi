package infra

import (
	"fmt"
	"webapi/pkg/app"
	httphandlers "webapi/pkg/interface/http_handlers"
	"webapi/pkg/presentation"

	"github.com/gin-gonic/gin"
)

type WebApiConfig struct {
	AppHost string
	AppPort int
}

func Start(config *WebApiConfig) error {
	router := gin.Default()

	// Register All Routes
	userService := app.NewUserService()
	userHTTPHandler := httphandlers.NewUserHTTPHandler(userService)
	presentation.NewUserRoutes(router, userHTTPHandler)

	return router.Run(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
