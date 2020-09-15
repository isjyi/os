package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isjyi/os/model"
)

type JWTManager struct {
	SecretKey string
	ExpiresAt time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	ID         uint   `json:"id"`
	NickName   string `json:"nick_name"`
	HeaderImg  string `json:"header_img"`
	BufferTime int64
}

func (manager *JWTManager) CreateToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.SecretKey))
}

func (manager *JWTManager) Generate(user model.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.ExpiresAt).Unix(),
			NotBefore: time.Now().Unix() - 1000,
			Issuer:    "os",
		},
		ID:         user.ID,
		NickName:   user.NickName,
		HeaderImg:  user.HeaderImg,
		BufferTime: 60 * 60 * 24,
	}

	return manager.CreateToken(claims)
}

func (manager *JWTManager) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.SecretKey), nil
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
