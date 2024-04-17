package config

import (
	"database/sql"
	"fmt"
)

func (r *Config) DatabaseOpen() error {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", r.Variable.dbUser, r.Variable.dbPass, r.Variable.dbHost, r.Variable.dbPort, r.Variable.dbName)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	r.Database = db

	return nil
}
