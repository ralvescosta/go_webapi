package env

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func ConfigEnvs() {
	profile := os.Getenv("GO_ENV")

	viper.SetConfigName(".env." + strings.ToLower(profile))
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("An error occurred reading the config file", err)
	}
}
