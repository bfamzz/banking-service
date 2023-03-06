package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error that can be returned by the VerifyToken function
var (
	ErrInvalidToken          = errors.New("token is invalid")
	ErrExpiredToken          = errors.New("token has expired")
	ErrInvalidOrExpiredToken = errors.New("token is either invalid or has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	NotBefore time.Time `json:"not_before"`
	ExpiresAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payuload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID.String(),
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}
