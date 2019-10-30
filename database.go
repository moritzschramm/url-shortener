package main

import ( 
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


const (

	DB_DSN = "links.db"				// sqlite database file
	DB_INIT_STMT = "init_db.sql"	// sql statement to create initial tables
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

	// read init statement from file DB_INIT_STMT
	initStmt, err := ioutil.ReadFile(DB_INIT_STMT)
	if err != nil {
		log.Fatal("Error reading database init statement file: ", err.Error())
	}

	// create tables (if not already exisiting)
	_, err = db.Exec(string(initStmt))
	if err != nil {
		log.Fatal("Error executing database init statement: ", err.Error())
	}

	return db
}