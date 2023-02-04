package utils

import (
	"fmt"
	"harry/query-overflow/models"
	"time"

	"github.com/golang-jwt/jwt"
)

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *models.SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("verysecure"), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}
