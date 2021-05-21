package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	privateKeyInBytes, _ := ioutil.ReadFile("cert/id_rsa")

	fmt.Println(privateKeyInBytes)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyInBytes)
	if err != nil {
		log.Fatal("privateKey", err)
	}

	claims := make(jwt.MapClaims)
	claims["jti"] = "1"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "iss"
	claims["aud"] = "aud"

	t, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(t)
}
