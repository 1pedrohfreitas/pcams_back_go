package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "pcam.postgres.uhserver.com"
	port     = 5432
	user     = "pedrohfreitas"
	password = "Pcam01@"
	dbname   = "pcam"
)

var dbSql *sql.DB

func StartDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	dbSql = db

	// close database
	// defer db.Close()

	// check db
	// err = db.Ping()
	// CheckError(err)

	// fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDataBase() *sql.DB {
	return dbSql
}
