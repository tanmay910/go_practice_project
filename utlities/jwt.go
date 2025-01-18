package utlities

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"errors"
	"fmt"
)

const secretKey = "secretKey"
func GenerateToken(email string , userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (int64,error) {
	parsedtoken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return  []byte(secretKey), nil
	})

	
	if err != nil {
		return  0,err
	}
	tokenISValid := parsedtoken.Valid
	if !tokenISValid {
		return 0,errors.New("token is invalid")
	}

	claims , ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("could not parse claims")
	}

	// email,_ := claims["email"].(string)
	var userId int64
	// Ensure the value is asserted as float64 before converting to int64
	if userIdFloat, ok := claims["userId"].(float64); ok {
		userId := int64(userIdFloat)
		return userId, nil
	} else {
		return 0, fmt.Errorf("invalid userId type")
	}


	return userId,nil

}
