package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("your_secret_key")

type Claims struct {
	UserID string `json:"id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrInvalidKey
	}
}

func IsTokenExpired(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, nil)
	if token == nil {
		return true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp := claims["exp"].(float64)
		expiryTime := time.Unix(int64(exp), 0)
		return expiryTime.Before(time.Now())
	}

	return true
}
