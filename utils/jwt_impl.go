package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretKey = []byte("your_secret_key")

type Claims struct {
	UserID   uuid.UUID `json:"id"`
	UserType string    `json:"userType"`
	ProfileID  uuid.UUID 
	jwt.StandardClaims
}

func GenerateAccessToken(userID uuid.UUID, userType string, profileId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID:   userID,
		UserType: userType,
		ProfileID:  profileId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute * 24).Unix(),
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
