package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Temporary signing of keys, to store in some other place not exposed to repository
const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // valid for 2 hours
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Type checking syntax to check if signed same
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token.")
	}
	//Check validity of token
	tokenValid := parsedToken.Valid
	if !tokenValid {
		return 0, errors.New("Invalid token!")
	}
	//Check invalid token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims!")
	}
	//No ok needed as we did it above
	//email := claims["email"].(string)
	userID := int64(claims["userId"].(float64))
	return userID, nil
}
