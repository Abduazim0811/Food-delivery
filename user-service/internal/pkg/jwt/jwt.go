package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Abdu0811")

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func CheckToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token expired")
			}
			if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, fmt.Errorf("invalid signature")
			}
		}
		return nil, fmt.Errorf("token parsing error: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
