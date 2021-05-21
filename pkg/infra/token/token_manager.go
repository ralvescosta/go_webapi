package token

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken() (string, error) {
	privateKeyInBytes, err := ioutil.ReadFile("cert/id_rsa")
	if err != nil {
		return "", errors.New("error when try to read rsa private key")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyInBytes)
	if err != nil {
		return "", errors.New("error when try to create rsa private key")
	}

	claims := jwt.StandardClaims{
		Audience:  "aud",
		Issuer:    "iss",
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		Id:        "1",
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Add(time.Hour * 1).Unix(),
	}

	t, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return "", errors.New("error when try to create jwt")
	}

	return t, nil
}

func VerifyToken(token string) (interface{}, error) {
	publicKeyInBytes, err := ioutil.ReadFile("cert/id_rsa.pub")
	if err != nil {
		return "", errors.New("error when try to read rsa public key")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyInBytes)
	if err != nil {
		return "", errors.New("error when try to create rsa public key")
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return "", errors.New("invalid token")
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims, nil
}
