package utils

import (
	"log"
	"os"
	"time"

	"harry/query-overflow/models"

	jwt "github.com/golang-jwt/jwt"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both the detailed token and in future refresh token too :)
func GenerateAllTokens(email string, name string, uid string) (signedToken string, err error) {
	claims := &models.SignedDetails{
		Email: email,
		Name:  name,
		Id:    uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}
