package storage

import (
	"fmt"

	"github.com/motorheads/auth_service/config"
	"github.com/motorheads/auth_service/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(data map[string]string) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

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
