package infra

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"webapi/pkg/app/services"
	httphandlers "webapi/pkg/handlers/http_handlers"
	"webapi/pkg/infra/hasher"
	"webapi/pkg/infra/repos"
	"webapi/pkg/infra/token"
	presenter "webapi/pkg/presenter"
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
	// Create Infra Instancies
	userRepo := repos.NewUserRepository(db)
	hasher := hasher.NewHahser()
	tokenManager := token.NewTokenManager()

	// Register All Routes
	userService := services.NewUserService(userRepo, hasher)
	userHTTPHandler := httphandlers.NewUserHTTPHandler(userService)
	presenter.NewUserRoutes(router, userHTTPHandler)

	authService := services.NewAuthenticationUser(userRepo, hasher, tokenManager)
	authHTTPHandler := httphandlers.NewAuthenticationHTTPHandler(authService)
	presenter.NewAuthenticationRoute(router, authHTTPHandler)

	return router.Run(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
