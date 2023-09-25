package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func ValidateJWT(currentToken string) error {
	if ContainsToken(currentToken) {
		if CachedTokens[currentToken].Before(time.Now()) {
			return jwt.ErrTokenExpired
		}
	}

	claims := &Claims{}
	tokenObject, err := jwt.ParseWithClaims(currentToken, claims, func(token *jwt.Token) (any, error) {
		return signingKey, nil
	})

	if err != nil {
		return err
	}

	if !tokenObject.Valid {
		return jwt.ErrTokenSignatureInvalid
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return jwt.ErrTokenExpired
	}

	CachedTokens[currentToken] = claims.ExpiresAt.Time
	return nil
}
