package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// ConnectToDB will establish a connection to the database.
func ConnectToDB(user, password, address, dbname string) {
	dab, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True", user, password, address, dbname))
	if err != nil {
		log.Fatalln(err)
	}
	dab.SetMaxOpenConns(5000)
	dab.SetMaxIdleConns(400)
	DB = dab
}

// CloseDB cleanly closes the DB connection.
func CloseDB() error {
	return DB.Close()
}
