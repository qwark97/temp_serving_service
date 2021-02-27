package main

import (
	"database/sql"
	"log"
)

func getCurrentTemp() (float32, string, error) {
	sqlStatement := selectStatements.getStatement("currentTemp")
	var temp float32
	var unit string
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&temp, &unit); err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		return 0, "unknown", nil
	case nil:
		return temp, unit, nil
	default:
		log.Println("ERROR - could not select data from DB - ", err)
		return 0, "unknown", err
	}
}