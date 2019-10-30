package main

import ( 
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


const (

	DB_DSN = "links.db"
)

func SetupDatabase() *sql.DB {

	// create database interface
	db, err := sql.Open("sqlite3", DB_DSN)
	if err != nil {
		log.Fatal("Error opening database: ", err.Error())
	}

	// check connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	// init database tables
	err = initTables(db, config)
	if err != nil {
		log.Fatal("Error executing database init statement: ", err.Error())
	}

	// read init statement from file DB_INIT_STMT
	initStmt, err := ioutil.ReadFile(DB_INIT_STMT)
	if err != nil {
		log.Fatal("Error reading database init statement file: ", err.Error())
	}

	// create tables (if not already present)
	_, err = db.Exec(string(initStmt))
	if err != nil {
		return err
	}

	return db
}