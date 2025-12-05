package storage

import (
	"context"
	"university_app/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) Authorize(login, password string) (*models.User, error) {
	query := `SELECT password, is_admin FROM users WHERE login = $1 LIMIT 1`
	var hash string
	var isAdmin bool
	_ = s.pool.QueryRow(context.Background(), query, login).Scan(&hash, &isAdmin)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, err
	}

	user := models.User{login, isAdmin}
	return &user, nil
}
