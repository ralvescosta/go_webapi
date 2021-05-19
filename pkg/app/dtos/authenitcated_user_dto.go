package dtos

import "time"

type AuthenticatedUserDto struct {
	AccessToken string
	ExpireIn    time.Time
}
