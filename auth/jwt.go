package auth

import (
	"fmt"
	jwtGo "github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(username string) (string, error) {
	claims := jwtGo.MapClaims{}
	claims["authorized"] = true
	claims["user_name"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ValidateToken(token string) (*jwtGo.Token, error) {
	parsedToken, parsedTokenErr := jwtGo.Parse(token, func(token *jwtGo.Token) (interface{}, error) {
		// Check if token signed
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method, %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if parsedTokenErr != nil {
		return nil, parsedTokenErr
	}
	return parsedToken, nil
}

func GetUserInformationFromToken(token string, key string) (interface{}, error) {
	parsedToken, parsedTokenErr := ValidateToken(token)
	if parsedTokenErr != nil {
		return "", parsedTokenErr
	}
	if claims, ok := parsedToken.Claims.(jwtGo.MapClaims); ok && parsedToken.Valid {
		if userInformation, existed := claims[key]; existed {
			return userInformation, nil
		}
	}
	return "", fmt.Errorf("token invalid, please try again")
}
