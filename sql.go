package main

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const insertQuery=`
insert into attendance (email, name, location, review)
values ($1, $2, $3, $4) 
`

func insertAttendee(conn *pgx.Conn, person Attendee) error {
	_, err := conn.Exec(context.Background(), insertQuery, person.Email, person.Name, person.Location, person.Review)
	if err != nil {
		return err
	}

	return nil
}