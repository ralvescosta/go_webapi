package infra

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"

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

	Env     string
	AppHost string
	AppPort int
	GinMode string

	WebApiReqsLog string
}

func Start(config *WebApiConfig) error {
	log.SetFormatter(&log.JSONFormatter{})
	standardFields := log.Fields{
		"hostname": config.AppHost,
		"appname":  "GoWebApi",
	}
	log.WithFields(standardFields)

	if config.Env != "dev" {
		writerReqsLogs, err := os.Create(config.WebApiReqsLog)
		if err != nil {
			err = fmt.Errorf("server.Start - create log writer")
			log.Fatal(err)
		}

		gin.DefaultWriter = io.MultiWriter(writerReqsLogs)
		log.SetOutput(writerReqsLogs)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(apmgin.Middleware(router))

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
