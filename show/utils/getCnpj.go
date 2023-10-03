package utils

import (
	"github.com/golang-jwt/jwt"
)

func GetCnpj(tokenString string) string {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		//add log
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["cnpj"].(string)
	}

	//add log
	return ""
}
