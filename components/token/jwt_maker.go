package token

import (
	"cid-10-api/models"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secretKey string
}

const minSecretKeySize = 32

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := models.NewPayload(username, duration)

	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
	
}

func (maker *JWTMaker) VerifyToken(token string) (*models.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, models.ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &models.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, models.ErrExpiredToken) {
			return nil, models.ErrExpiredToken
		}

		return nil, models.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*models.Payload)
	if !ok {
		return nil, models.ErrInvalidToken
	}

	return payload, nil
 }

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}