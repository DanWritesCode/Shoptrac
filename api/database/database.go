package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

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

func SetDatabaseConfig(key, value string) error {
	results, err := DB.Query("SELECT COUNT(id) FROM config WHERE name = ?;", key)
	if err != nil {
		return err
	}

	for results.Next() {
		count := 0
		err = results.Scan(&count)
		if err != nil {
			return err
		}

		if count > 0 {
			_, err = DB.Exec("UPDATE config SET `value` = ? WHERE `name` = ?;", value, key)
			if err != nil {
				return err
			}
		} else {
			_, err = DB.Exec("INSERT INTO `config` (`id`, `name`, `value`) VALUES (NULL, ?, ?);", key, value)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// CloseDB cleanly closes the DB connection.
func CloseDB() error {
	return DB.Close()
}
