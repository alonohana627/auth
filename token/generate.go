package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(email string, nickname string) (string, error) {
	payload, err := CreatePayload(email, nickname, time.Duration(200)*time.Hour)
	if err != nil {
		return "", err
	}
	claims := &Claims{
		Email:    email,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(signingKey)
	if err != nil {
		CachedTokens[token] = claims.ExpiresAt.Time
	}
	return token, err
}
