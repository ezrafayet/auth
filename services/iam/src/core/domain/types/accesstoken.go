package types

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type AccessToken string

func NewAccessToken(userId Id, role string) (AccessToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	ts := time.Now().UTC()

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "iam"
	claims["aud"] = "iam"
	claims["exp"] = ts.Add(time.Minute * 15).Unix()
	claims["user_id"] = string(userId)
	claims["role"] = role
	claims["server_region"] = os.Getenv("REGION")
	claims["issued_at"] = ts.Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_PRIVATE_KEY")))

	if err != nil {
		return "", err
	}

	return AccessToken(tokenString), nil
}

// func ParseAndValidateAccessToken (str string) (AccessToken, error) {
//
// }
