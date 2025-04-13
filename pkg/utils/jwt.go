package utils

import (
	"errors"
	"golang-mongo-auth/pkg/common/messages"
	"golang-mongo-auth/pkg/common/types"
	"golang-mongo-auth/pkg/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetJwtSecrets())

type Claims struct {
	UserId string     `json:"userId"`
	Role   types.Role `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(id string, role types.Role) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (*Claims, error) {

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.New(messages.ErrMissingToken)
	}

	claims, err := ParseToken(tokenString)

	if err != nil {
		return nil, errors.New(messages.ErrInvalidToken)
	}

	return claims, nil
}
