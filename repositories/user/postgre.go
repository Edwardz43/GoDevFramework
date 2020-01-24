package user

import (
	"Edwardz43/godevframework/models"
	"database/sql"
	"log"
)

// PostgreUserRepository ...
type PostgreUserRepository struct {
	db *sql.DB
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// GetInstance ...
func GetInstance(db *sql.DB) *PostgreUserRepository {
	return &PostgreUserRepository{db: db}
}

// Create ..
func (r *PostgreUserRepository) Create(name, email string) (bool, error) {
	statement := "INSERT INTO public.\"User\" (name, email) VALUES ($1, $2);"
	_, err := r.db.Exec(statement, name, email)
	if err != nil {
		return false, err
	}
	return true, nil
}

// FindOne ..
func (r *PostgreUserRepository) FindOne(id string) (*models.User, error) {
	statement := "SELECT * FROM public.\"User\" WHERE id=$1;"

	rows, err := r.db.Query(statement, id)

	if err != nil {
		return nil, err
	}

	u := models.User{}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.Email)
		checkErr(err)
	}

	return &u, nil
}

// Update ..
func (r *PostgreUserRepository) Update(id, name, email string) (bool, error) {
	statement := "UPDATE public.\"User\" SET name=$1, email=$2 WHERE id=$3;"

	result, err := r.db.Exec(statement, name, email, id)
	if err != nil {
		return false, err
	}

	i, err := result.RowsAffected()

	return i == 1, nil
}
