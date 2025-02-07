package jwtToken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwtAndGetData(jwtToken string, key string) (interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodES256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	data, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token verification failed")
	}

	if exp, ok := data["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			return data, fmt.Errorf("Token has expired at: %v", expirationTime)
		}
	}

	return data, nil
}

func CreateNewToken(data interface{}, key string) (string, error) {
	claims := jwt.MapClaims(data)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
