package token

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/interfaces"
)

type tokenManager struct{}

func (t tokenManager) GenerateToken(tokenData *dtos.TokenDataDto) (string, error) {
	privateKeyInBytes, err := ioutil.ReadFile("cert/id_rsa")
	if err != nil {
		log.Println(err)
		return "", errors.New("error when try to read rsa private key")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyInBytes)
	if err != nil {
		log.Println(err)
		return "", errors.New("error when try to create rsa private key")
	}

	claims := jwt.StandardClaims{
		Audience:  tokenData.Audience,
		Issuer:    "Go WebApi Templete",
		ExpiresAt: tokenData.ExpireIn.Unix(),
		Id:        fmt.Sprintf("%d", tokenData.Id),
		IssuedAt:  time.Now().Unix(),
		NotBefore: tokenData.ExpireIn.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		log.Println(err)
		return "", errors.New("error when try to create jwt")
	}

	return token, nil
}

func (t tokenManager) VerifyToken(token string) (*dtos.AuthenticatedUserDto, error) {
	publicKeyInBytes, err := ioutil.ReadFile("cert/id_rsa.pub")
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to read rsa public key")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyInBytes)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to create rsa public key")
	}

	tok, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			log.Println(err)
			return nil, errors.New("unexpected method")
		}
		return publicKey, nil
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid token")
	}

	claims, ok := tok.Claims.(*jwt.StandardClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("jwt is expired")
	}

	id, _ := strconv.ParseUint(claims.Id, 10, 32)
	return &dtos.AuthenticatedUserDto{
		Id:          id,
		AccessToken: token,
		ExpireIn:    time.Unix(claims.ExpiresAt, 0),
	}, nil
}

func NewTokenManager() interfaces.ITokenManager {
	return &tokenManager{}
}
