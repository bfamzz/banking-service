package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecurityKeySize = 32

// JWTMaker is a JSON Web Token Maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecurityKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecurityKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        payload.ID,
		Subject:   payload.Username,
		IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		NotBefore: jwt.NewNumericDate(payload.NotBefore),
		ExpiresAt: jwt.NewNumericDate(payload.ExpiresAt),
	})
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidOrExpiredToken
	}

	claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &Payload{
		ID:        claims.ID,
		Username:  claims.Subject,
		IssuedAt:  claims.IssuedAt.Time,
		NotBefore: claims.NotBefore.Time,
		ExpiresAt: claims.ExpiresAt.Time,
	}, nil
}
