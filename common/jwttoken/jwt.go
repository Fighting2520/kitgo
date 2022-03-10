package jwttoken

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var ErrJwtTokenExpired = errors.New("toke expired")

func GenerateToken(secret, username, password string, expireSeconds int) (string, error) {
	expireTime := time.Now().Add(time.Duration(expireSeconds) * time.Second)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      expireTime.Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func parseToken(token, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

func ParseUsernameAndPasswordFromToken(token, secret string) (string, string, error) {
	jwtToken, err := parseToken(token, secret)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", "", ErrJwtTokenExpired
			}
		}
		return "", "", err
	}
	if !jwtToken.Valid {
		return "", "", errors.New("invalid jwt token")
	}
	var claims = jwtToken.Claims.(jwt.MapClaims)
	var username, password string
	if v, ok := claims["username"]; ok {
		username = v.(string)
	}
	if v, ok := claims["password"]; ok {
		password = v.(string)
	}
	return username, password, nil
}
