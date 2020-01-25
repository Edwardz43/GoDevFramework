package migrate

import (
	"Edwardz43/godevframework/config"
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgresMigration ...
type PostgresMigration struct {
}

// GetPostgresInstance return a postgresSQL migration instance.
func GetPostgresInstance() *PostgresMigration {
	return &PostgresMigration{}
}

// Migrate migrates database.
func (m *PostgresMigration) Migrate() (bool, error) {
	connStr := config.Database()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false, err
	}

	defer db.Close()

	_, err = db.Exec(
		`Create table IF NOT EXISTS "User" (
			id serial,
			name varchar(20),
			email varchar(50)
		);`)
	if err != nil {
		return false, err
	}
	return true, nil
}
