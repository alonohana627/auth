package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

var signingKey []byte

type Payload struct {
	ID               uuid.UUID
	Email            string
	Nickname         string
	IssuedAt         time.Time
	ExpiredAt        time.Time
	RefreshExpiresAt time.Time
}

type Claims struct {
	Email    string
	Nickname string
	jwt.RegisteredClaims
}

func CreatePayload(email string, nickname string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	payload := &Payload{
		ID:               tokenID,
		Email:            email,
		Nickname:         nickname,
		IssuedAt:         now,
		ExpiredAt:        now.Add(duration),
		RefreshExpiresAt: now.Add(duration),
	}

	return payload, nil
}

func init() {
	signingKey = []byte(os.Getenv("JWT_SECRET"))
	CachedTokens = make(map[string]time.Time)
}
