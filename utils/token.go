package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type JwtUserClaims struct {
	jwt.StandardClaims
	UserId *uuid.UUID `json:"user_id"`
}

func CreateJWTUser(ID *uuid.UUID) (string, error) {
	tokenDuration, _ := time.ParseDuration(viper.GetString("token.lifeTime"))
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["ExpiresAt"] = time.Now().Add(tokenDuration).Unix()
	claims["UserId"] = ID

	signedToken, err := token.SignedString([]byte(viper.GetString("token.secretKeyJWT")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
