package jwtauth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateRefreshToken(id uint) (string, error) {

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id,
		"name": "refresh",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "error", err
	}

	return tokenString, nil
}

func CreateAccessToken(id uint) (string, error) {

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id,
		"exp":  time.Now().Add(time.Hour * 6).Unix(),
		"name": "access",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "error", err
	}

	return tokenString, nil
}

func DecodeAccessToken(tokenString string) (uint, error) {
	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g., []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return 0, err // Return the error
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// Check the exp
		expirationTime, ok := claims["exp"].(float64)
		if !ok {
			return 0, jwt.ErrTokenInvalidClaims
		}

		if float64(time.Now().Unix()) > expirationTime {
			return 0, jwt.ErrTokenExpired
		}

		// Token is valid, return the subject (sub) claim
		subject, ok := claims["sub"]
		if !ok {
			return 0, jwt.ErrTokenInvalidId
		}
		sub := uint(subject.(float64))

		return sub, nil
	}

	return 0, jwt.ErrTokenSignatureInvalid
}

func DecodeRefreshToken(tokenString string) (uint, error) {
	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g., []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return 0, err // Return the error
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// Token is valid, return the subject (sub) claim
		subject, ok := claims["sub"]
		if !ok {
			return 0, jwt.ErrTokenInvalidId
		}
		sub := uint(subject.(float64))

		return sub, nil
	}

	return 0, jwt.ErrTokenSignatureInvalid
}
