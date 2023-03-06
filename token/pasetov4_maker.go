package token

import (
	"fmt"
	"log"
	"time"

	"aidanwoods.dev/go-paseto"
	"golang.org/x/crypto/chacha20"
)

// PasetoMaker is a PASETO token maker
type PasetoV4Maker struct {
	paseto       paseto.Token
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoV4Maker(secretKey string) (Maker, error) {
	if len(secretKey) != chacha20.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20.KeySize)
	}

	symmetricKey, err := paseto.V4SymmetricKeyFromBytes([]byte(secretKey))
	if err != nil {
		log.Fatal("cannot generate symmetric key from bytes:", err)
	}

	maker := &PasetoV4Maker{
		paseto:       paseto.NewToken(),
		symmetricKey: symmetricKey,
	}

	return maker, nil
}

func (maker *PasetoV4Maker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	maker.paseto.SetString("id", payload.ID)
	maker.paseto.SetSubject(payload.Username)
	maker.paseto.SetIssuedAt(payload.IssuedAt)
	maker.paseto.SetNotBefore(payload.NotBefore)
	maker.paseto.SetExpiration(payload.ExpiresAt)

	return maker.paseto.V4Encrypt(maker.symmetricKey, nil), payload, nil
}

func (maker *PasetoV4Maker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	pasetoToken, err := parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		return nil, ErrInvalidOrExpiredToken
	}

	id, err := pasetoToken.GetString("id")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	subject, err := pasetoToken.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	issuedAt, err := pasetoToken.GetIssuedAt()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	notBefore, err := pasetoToken.GetNotBefore()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	expiresAt, err := pasetoToken.GetExpiration()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Payload{
		ID:        id,
		Username:  subject,
		IssuedAt:  issuedAt,
		NotBefore: notBefore,
		ExpiresAt: expiresAt,
	}, nil
}
