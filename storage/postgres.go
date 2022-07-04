package storage

import (
	"fmt"

	"github.com/motorheads/auth_service/config"
	"github.com/motorheads/auth_service/models"
)

func CreateUser(user models.User) error {
	query := `
		INSERT INTO users(
			name,
			email,
			password
		) VALUES (
			$1,
			$2,
			$3
		);
	`

	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf(`
		SELECT * 
		FROM users
		WHRER email=%s
	`, email)

	row := config.DB.QueryRow(query)

	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	return &user, err
}

func GetUserById(id string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf(`
		SELECT * 
		FROM users
		WHRER id=%s
	`, id)

	row := config.DB.QueryRow(query)

	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	return &user, err
}
