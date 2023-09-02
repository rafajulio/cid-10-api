package token

import (
	"cid-10-api/models"
	"time"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*models.Payload, error)
}