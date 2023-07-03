package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("user"),
			os.Getenv("pass"),
			os.Getenv("hostdb"),
			os.Getenv("port"),
			os.Getenv("db_name"))
	)
	fmt.Print(ConnectionString)

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal("Error opening database")
	}
	
	fmt.Print(db)

}