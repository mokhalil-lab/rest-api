package utils

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string)(string, error){
	hash , err := bcrypt.GenerateFromPassword([]byte(password) , 14)
	return string(hash) , err
}

func Comparehashwithinput( hashedpassword , password string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword) , []byte(password))
	return err == nil
}

