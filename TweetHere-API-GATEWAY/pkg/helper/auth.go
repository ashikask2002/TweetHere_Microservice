package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func ValidateToken(tokenString string) (*authCustomClaimsAdmin, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsAdmin); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
