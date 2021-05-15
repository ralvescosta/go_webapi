package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func ConfigEnvs() {
	profile := os.Getenv("SERVICES_PROFILE")

	viper.SetConfigName(".env." + strings.ToLower(profile))
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("An error occurred reading the config file", err)
	}

	fmt.Println(viper.GetString("db_host"))
}
