package repositories

import (
	"database/sql"

	"example.com/postgresdatabase/database/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *models.User) error {

	query := `
    INSERT INTO users (name, email)
    VALUES ($1, $2)
    RETURNING id
    `

	return r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
	).Scan(&user.ID)
}
