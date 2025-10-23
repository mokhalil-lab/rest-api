package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID int64
	Email string
	Password string
}

func (u *User) Save() error {
	query := `INSERT INTO users(email , password) VALUES(?,?)`
	stmt , err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedpassword , err := utils.Hashpassword(u.Password)
	if err != nil {
		return err
	}
	result , err := stmt.Exec(u.Email , hashedpassword)
	if err != nil {
		return err
	}
	id , err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (u *User) Authenticate() error {
	query := `SELECT id , password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query , u.Email)
	var retrievedpassword string
	err := row.Scan(&u.ID , &retrievedpassword)
	if err != nil {
		return errors.New("couldn't authenticate user")
	}
	isvalid := utils.Comparehashwithinput(retrievedpassword , u.Password)
	if !isvalid {
		return errors.New("couldn't authenticate user")
	}
	return nil
}