package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do

type jwtCustomClaim struct {
	USER_ID string
	ROLE_ID int
	jwt.StandardClaims
}

type jwtService struct{}

//NewJWTService method is creates a new instance of JWTService

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "auth12&6%(?>,g#"
	}
	return secretKey
}

func GenerateToken(USER_ID string, ROLE_ID int) string {
	claims := &jwtCustomClaim{
		USER_ID,
		ROLE_ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 1).Unix(),
			Issuer:    "Auth",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(getSecretKey()))

	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(getSecretKey()), nil
	})
}

// get userid from token
func GetUserID(authHeader string) uint {
	token, errToken := ValidateTokenVal(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["USER_ID"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return uint(id)
}

func ValidateTokenVal(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(getSecretKey()), nil
	})
}
