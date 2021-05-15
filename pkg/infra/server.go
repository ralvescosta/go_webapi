package infra

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"webapi/pkg/app/services"
	httphandlers "webapi/pkg/interfaces/http_handlers"
	"webapi/pkg/presentation"
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
		return errors.New("DB CONNECTION ERROR - " + err.Error())
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// Register All Routes
	userService := services.NewUserService()
	userHTTPHandler := httphandlers.NewUserHTTPHandler(userService)
	presentation.NewUserRoutes(router, userHTTPHandler)

	return router.Run(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
