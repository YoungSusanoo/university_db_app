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
	_ = s.pool.QueryRow(context.Background(), query, login).Scan(&hash, &isAdmin)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, err
	}

	user := models.User{login, isAdmin}
	return &user, nil
}

func (s *Storage) GetSubjects() ([]models.Subject, error) {
	query := `SELECT * FROM subjects`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Subject])
}

func (s *Storage) GetStudents() ([]models.Student, error) {
	query := `SELECT id, first_name, last_name, father_name, group_id FROM people WHERE type = 'S'`
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
