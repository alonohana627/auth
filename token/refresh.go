package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func RefreshJWT(currentToken string) (string, error) {
	valStatus := ValidateJWT(currentToken)
	if valStatus != nil {
		return "", valStatus
	}

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(currentToken, claims, func(token *jwt.Token) (any, error) {
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(200 * time.Hour)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := jwtToken.SignedString(signingKey)

	if err != nil {
		CachedTokens[newToken] = expirationTime
	}
	return newToken, err
}
