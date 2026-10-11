package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	Role string `json:"role"`
}

func CreateToken(jwtSecret string, userID string, role string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID,
		},
		Role: role,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(jwtSecret string, tokenString string) {
	var claims Claims

	parsed, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signin method :v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return Claims{}, fmt.Errorf("parsed token failed %w", err)
	}

	if !parsed.Valid {
		return Claims{}, errors.New("invalid token")
	}

	if claims.Subject == "" {
		return claims{}, errors.New("token missing subject")
	}

	return claims, nil

}
