package driver

import (
	"database/sql"
	"fmt"
	"os"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
