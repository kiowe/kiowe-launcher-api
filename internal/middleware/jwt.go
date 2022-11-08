package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateNewAccessToken(id int, devPubAcc bool) (string, error) {
	claims := jwt.MapClaims{
		"id":          id,
		"dev_pub_acc": devPubAcc,
		"exp":         time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
