package security

import (
	"backend-github-trending/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SECRET_KEY = "thien"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClains{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
