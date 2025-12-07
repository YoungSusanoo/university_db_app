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

func (s *Storage) GetStudents() ([]models.Student, error) {
	query := `SELECT people.id, first_name, last_name, father_name, g.name FROM people, groups g WHERE type = 'S' AND g.id = group_id`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Student])
}

func (s *Storage) GetTeachers() ([]models.Teacher, error) {
	query := `SELECT id, first_name, last_name, father_name FROM people WHERE type = 'P'`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Teacher])
}
