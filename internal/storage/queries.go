package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) Authorize(login, password string) (*models.User, error) {
	query := `SELECT password, is_admin FROM users WHERE login = $1 LIMIT 1`
	var hash string
	var isAdmin bool
	err := s.pool.QueryRow(context.Background(), query, login).Scan(&hash, &isAdmin)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, err
	}

	user := models.User{Name: login, IsAdmin: isAdmin}
	return &user, nil
}

func (s *Storage) GetStudentGroupRelation() (map[string]int64, error) {
	query := `SELECT name, id FROM groups`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var key string
	var value int64
	nameId := make(map[string]int64)
	_, err = pgx.ForEachRow(rows, []any{&key, &value}, func() error {
		nameId[key] = value
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nameId, nil
}
