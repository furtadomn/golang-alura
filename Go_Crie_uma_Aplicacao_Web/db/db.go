package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}

	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	pass := os.Getenv("PASSWORD")

	connect := ("user=" + user + " dbname=" + dbname + " password=" + pass + " host=localhost sslmode=disable")

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err.Error())
	}

	return db
}
