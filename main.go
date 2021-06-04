package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"webapi/cmd"
)

func main() {
	err := cmd.WebApi()
	if err != nil {
		err = fmt.Errorf("error while start application: %v", err)
		log.Fatal(err)
	}
}
