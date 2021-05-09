package infra

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type WebApiConfig struct {
	AppHost string
	AppPort int
}

func Start(config *WebApiConfig) error {
	router := gin.Default()

	return router.Run(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
