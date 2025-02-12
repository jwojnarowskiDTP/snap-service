package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "pg-pass"
	dbname   = "postgres"
)

func main() {
	// connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// connection to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// connection check
	err = db.Ping()
	if err != nil {
		log.Fatalf("Problem with database ping: %v", err)
	}
	fmt.Println("Database connection sucesful!")

	// create user rizone
	createUserQuery := `CREATE USER rizone WITH LOGIN PASSWORD 'rizone';`
	_, err = db.Exec(createUserQuery)
	if err != nil {
		log.Printf("Database user init failure: %v", err)
	} else {
		fmt.Println("Database user 'rizone' created.")
	}

	// create database rizone
	createDbQuery := `CREATE DATABASE rizone OWNER rizone;`
	_, err = db.Exec(createDbQuery)
	if err != nil {
		log.Printf("Database init failure: %v", err)
	} else {
		fmt.Println("Database 'rizone' created.")
	}
}
