package infra

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

	router := gin.New()

	if config.Env != "dev" {
		writerReqsLogs, err := os.Create(config.WebApiReqsLog)
		if err != nil {
			err = fmt.Errorf("server.Start - create log writer")
			log.Fatal(err)
		}

		gin.DefaultWriter = io.MultiWriter(writerReqsLogs)

		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(writerReqsLogs)

		router.Use(LoggerToFile())
	} else {
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())
	router.Use(apmgin.Middleware(router))

	sqlDb, err := GetSQLConnection(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		return errors.New("DB CONNECTION ERROR - " + err.Error())
	}
	defer func() {
		if sqlDb != nil {
			sqlDb.Close()
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:123456@localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Create Infra Instancies
	userRepo := repos.NewUserSQLRepository(sqlDb)
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

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		log.WithContext(c).WithFields(
			log.Fields{
				"statusCode": c.Writer.Status(),
				"latency":    latencyTime.String(),
				"clientIP":   c.ClientIP(),
				"method":     c.Request.Method,
				"uri":        c.Request.RequestURI,
			},
		).Info("Request Log")
	}
}
