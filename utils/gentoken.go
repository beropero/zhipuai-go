package utils

import (
	"zhipuai-go/consts"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成鉴权token
func GetTocken() (token string, err error){
	expHours := 300 // Token expires in 300 Hours

	token, err = generateToken(consts.ApiKey, expHours)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateToken(apiKey string, expSeconds int) (string, error) {
	id, secret, found := splitAPIKey(apiKey)
	if !found {
		return "", fmt.Errorf("invalid apikey")
	}

	payload := jwt.MapClaims{
		"api_key":   id,
		"exp":       time.Now().Add(time.Hour*time.Duration(expSeconds)).UnixMilli(),
		"timestamp": time.Now().UnixMilli(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token.Header["sign_type"] = "SIGN"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func splitAPIKey(apiKey string) (id string, secret string, found bool) {
	parts := strings.Split(apiKey, ".")
	if len(parts) == 2 {
		return parts[0], parts[1], true
	}
	return "", "", false
}



