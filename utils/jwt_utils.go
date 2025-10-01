package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// generate token
func GenerateToken(userID uint, secretKey string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expiration time
			ID:        uuid.New().String(),                                // Unique identifier for the token
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Token issuance time
			NotBefore: jwt.NewNumericDate(time.Now()),                     // Token valid from
			Issuer:    "go-study-blog",                                    // Token issuer
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// parse token
func ParseToken(tokenString string, secretKey string) (*Claims, error) {

	claims := Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return &claims, nil

}
