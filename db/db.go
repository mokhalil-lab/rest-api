package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initdb(){
	var err error
	DB , err = sql.Open("sqlite3" , "api.db")
	if err != nil {
		log.Fatal("connection to db failed")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createtable()
}

func createtable(){
	createtableusers := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_ , err := DB.Exec(createtableusers)
	if err != nil {
		log.Fatal("couldn't create user table")
	}

	ceatetableevents := `

	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	date DATETIME NOT NULL,
	userid INTEGER,
	FOREIGN KEY(userid) REFERENCES users(id)
	)
	`
	_ , err = DB.Exec(ceatetableevents)
	if err != nil {
		log.Fatal("couldn't create event table")
	}

	createtableregistration := `
	CREATE TABLE IF NOT EXISTS registrations(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_ , err = DB.Exec(createtableregistration)
	if err != nil {
		log.Fatal("couldn't create register table")
	}
}