package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID int64
	Name string
	Description string
	Location string
	Date time.Time
	Userid int
}



func (e *Event) Save() error{
	query := `INSERT INTO events(name , description , location , date , userid)
	values(?,?,?,?,?)
	`
	stmt , err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result , err := stmt.Exec(e.Name ,e.Description , e.Location , e.Date , e.Userid)
	if err != nil {
		return err
	}

	id , err := result.LastInsertId()
	if err != nil {
		return err 
	}

	e.ID = id
	return nil

}

func Getallevents()([]Event,error){
	query := "SELECT * FROM events"
	rows , err := db.DB.Query(query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID , &event.Name , &event.Description , &event.Location , &event.Date , &event.Userid )
		if err != nil {
			return nil , err
		}
		events = append(events, event)
	}
	return events , nil 
}

func Geteventbyid(id int64)(*Event , error){
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query , id)
	var event Event
	err := row.Scan(&event.ID , &event.Name , &event.Description , &event.Location , &event.Date , &event.Userid)
	if err != nil {
		return nil , err 
	}
	return &event , nil
}

func (e Event)Update()error{
	query := `UPDATE events
	set name = ? , description = ? , location = ? , date = ?
	where id = ?
	`
	stmt , err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	
	_ , err = stmt.Exec(e.Name , e.Description , e.Location , e.Date , e.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete()error{
	query := `DELETE FROM events WHERE id = ?`
	stmt , err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_ , err = stmt.Exec(e.ID)
	if err != nil {
		return err 
	}
	return nil
}