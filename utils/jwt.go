package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "e82b9e67957e6f3eee7b3cbb19779644"

func Generatetoken(email string , id int64)(string , error){
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"email":email,
		"id": id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}


func Verifytoken(token string)(int64 , error){
	parsedtoken , err := jwt.Parse(token , func(token *jwt.Token)(interface{} ,error){
		_ , ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return 0 , errors.New("wrong method used")
		}

		return []byte(secretkey) , nil 
	})

	if err != nil {
		return 0 , errors.New("couldn't parse token")
	}

	isvalid := parsedtoken.Valid

	if !isvalid {
		return 0 , errors.New("token is invalid")
	}
	claims , ok := parsedtoken.Claims.(jwt.MapClaims)

	if !ok {
		return 0 , errors.New("not matching claims")
	}

	userid := int64(claims["id"].(float64))

	return userid , nil
}