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
			id serial NOT NULL,
			name varchar(20) NOT NULL,
			email varchar(50) NOT NULL,
			CONSTRAINT "User_pkey" PRIMARY KEY (id)
		);`)
	if err != nil {
		return false, err
	}
	return true, nil
}
