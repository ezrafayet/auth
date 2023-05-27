package types

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"iam/pkg/apperrors"
	"os"
	"time"
)

type AccessToken string

type CustomClaims struct {
	UserId       string `json:"user_id"`
	Roles        string `json:"roles"`
	ServerRegion string `json:"server_region"`
}

func NewAccessToken(customClaims CustomClaims, issuedAt time.Time) (AccessToken, int64, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expiresAt := issuedAt.Add(time.Minute * 15).Unix()

	standardClaims := jwt.StandardClaims{
		Audience:  "iam",
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt.Unix(),
		Issuer:    "iam",
	}

	claims := token.Claims.(jwt.MapClaims)

	claims["aud"] = standardClaims.Audience

	claims["exp"] = standardClaims.ExpiresAt

	claims["issued_at"] = standardClaims.IssuedAt

	claims["iss"] = standardClaims.Issuer

	claims["user_id"] = customClaims.UserId

	claims["roles"] = customClaims.Roles

	claims["server_region"] = customClaims.ServerRegion

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_PRIVATE_KEY")))

	if err != nil {
		return "", -1, err
	}

	return AccessToken(tokenString), expiresAt, nil
}

func ParseAndValidateAccessToken(tokenString string) (bool, CustomClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(apperrors.InvalidAccessToken)
		}
		return []byte(os.Getenv("JWT_PRIVATE_KEY")), nil
	})

	if err != nil {
		return false, CustomClaims{}, err
	}

	var customClaims CustomClaims

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		customClaims.UserId = claims["user_id"].(string)
		customClaims.Roles = claims["roles"].(string)
		customClaims.ServerRegion = claims["server_region"].(string)
	}

	return token.Valid, customClaims, nil
}
