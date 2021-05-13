package infra

import (
	"fmt"
	"webapi/pkg/app"
	httphandlers "webapi/pkg/interface/http_handlers"
	"webapi/pkg/presentation"

	"github.com/gin-gonic/gin"
)

type WebApiConfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	AppHost string
	AppPort int
}

func Start(config *WebApiConfig) error {
	router := gin.Default()

	db, err := GetConnection(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		return err
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// Register All Routes
	userService := app.NewUserService()
	userHTTPHandler := httphandlers.NewUserHTTPHandler(userService)
	presentation.NewUserRoutes(router, userHTTPHandler)

	return router.Run(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
