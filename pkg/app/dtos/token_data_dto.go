package dtos

import "time"

type TokenDataDto struct {
	Id       uint32
	ExpireIn time.Time
	Audience string
}
