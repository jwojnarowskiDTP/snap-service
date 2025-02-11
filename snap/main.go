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
	// Tworzenie connection stringa
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Nawiązywanie połączenia z bazą danych
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}
	defer db.Close()

	// Sprawdzenie połączenia
	err = db.Ping()
	if err != nil {
		log.Fatalf("Nie udało się pingować bazy danych: %v", err)
	}
	fmt.Println("Połączono z bazą danych!")

	// Tworzenie użytkownika rizone
	createUserQuery := `CREATE USER rizone WITH LOGIN PASSWORD 'rizone';`
	_, err = db.Exec(createUserQuery)
	if err != nil {
		log.Printf("Nie udało się utworzyć użytkownika: %v", err)
	} else {
		fmt.Println("Użytkownik 'rizone' został utworzony.")
	}

	// Tworzenie bazy danych rizone
	createDbQuery := `CREATE DATABASE rizone OWNER rizone;`
	_, err = db.Exec(createDbQuery)
	if err != nil {
		log.Printf("Nie udało się utworzyć bazy danych: %v", err)
	} else {
		fmt.Println("Baza danych 'rizone' została utworzona.")
	}
}
