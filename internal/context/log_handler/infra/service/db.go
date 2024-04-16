package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetDb() (*sql.DB, error) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get the current file's path")
	}

	dir := filepath.Dir(filename)

	err := godotenv.Load(dir + "/../../../../../.env")

	if err != nil {
		log.Fatal(err)
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
