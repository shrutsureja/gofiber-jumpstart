package utils

import (
	"app/src/config"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateSignedJWT(customClaims map[string]interface{}) (string, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(customClaims))

	return token.SignedString([]byte(cfg.JwtSecret))
}

func VerifySignedJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	cfg, err := config.GetConfig()

	if err != nil {
		return nil, nil, err
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JwtSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return token, nil, nil
}
