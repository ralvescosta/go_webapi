package main

import (
	"log"

	"webapi/cmd"
)

func main() {
	err := cmd.WebApi()
	if err != nil {
		log.Fatal("Deu ruim: ", err)
	}
}
