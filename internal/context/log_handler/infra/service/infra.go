package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDb() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
