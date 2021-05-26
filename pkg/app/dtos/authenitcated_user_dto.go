package dtos

import "time"

type AuthenticatedUserDto struct {
	Id          uint64
	AccessToken string
	ExpireIn    time.Time
}
