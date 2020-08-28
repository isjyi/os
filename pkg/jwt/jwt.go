package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isjyi/os/model"
	"github.com/spf13/viper"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
}

var Jwt *JWTManager

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

func (manager *JWTManager) Generate(user *model.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		ID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("server.jwtSecret")), nil
	})

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
		// fmt.Printf("%v %v", claims.ID, claims.StandardClaims.ExpiresAt)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("Timing is everything")
		} else {
			return nil, errors.New(fmt.Sprint("Couldn't handle this token:", err))
		}
	}
	return nil, errors.New("unknown mistake")
}
