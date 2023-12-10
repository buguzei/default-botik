package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Open() (*sql.DB, error) {
	database, err := sql.Open("postgres", "host=localhost port=5432 user=buguzei password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}
