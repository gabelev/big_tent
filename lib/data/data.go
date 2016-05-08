package data

import (
	"database/sql"
	"fmt"

	"github.com/gabelev/get_heard/lib/settings"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type DataStore interface {
	DbHealthCheck() (Status string, err error)
}

func New(db *sql.DB) (DB, error) {
	err := db.Ping()
	return DB{db}, err
}

func (db DB) DbHealthCheck() (Status string, err error) {
	query := fmt.Sprintf(
		"SELECT datname FROM pg_database WHERE datistemplate = false AND datname = '%s';", settings.Config.DbName,
	)
	err = db.QueryRow(query).Scan(&Status)
	return Status, err
}
