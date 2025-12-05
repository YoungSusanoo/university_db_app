package storage

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) HashUser(login, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`

	_, err = s.pool.Query(context.Background(), query, login, string(hash))
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Authorize(login, password string) (bool, error) {
	query := `SELECT password FROM users WHERE login = $1 LIMIT 1`
	var hash string
	_ = s.pool.QueryRow(context.Background(), query, login).Scan(&hash)

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	if err != nil {
		return false, nil
	}
	return true, nil
}
