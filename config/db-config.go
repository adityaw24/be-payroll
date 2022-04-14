package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Creating a new connection with postgresDB
func SetupDatabaseConnection() *sql.DB {
	// Load from .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// open the connection to db
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	CheckError(err)

	// close database
	//defer db.Close()

	// check the connection
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected to the database!")
	// return the connection
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
