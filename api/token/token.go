package token

import (
	"api_gateway/config"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ExtractClaimsAccess(tokenstr string) (jwt.MapClaims, error) {
	cfg := config.Load()
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.ACCES_TOKEN), nil
	})

	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		fmt.Println("Error getting")
		return nil, fmt.Errorf("token is invalid or not valid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token claims could not be converted to map claims")
	}
	return claims, nil

}

func ValidateToken(tokenstr string) (bool, error) {
	_, err := ExtractClaimsAccess(tokenstr)
	if err != nil {

		return false, err
	}

	return true, nil
}
