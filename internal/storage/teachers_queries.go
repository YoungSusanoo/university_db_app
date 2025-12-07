package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetTeachers() ([]models.Teacher, error) {
	query := `SELECT id, first_name, last_name, father_name FROM people WHERE type = 'P'`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Teacher])
}

func (s *Storage) InsertTeacher(teacher models.Teacher) error {
	query := `INSERT INTO people (first_name, last_name, father_name, type) VALUES($1, $2, $3, $4)`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		teacher.FirstName,
		teacher.LastName,
		teacher.FatherName,
		"P",
	)
	return err
}

func (s *Storage) UpdateTeacher(teacherOld, teacherNew models.Teacher) error {
	query := `UPDATE people SET first_name = $1, last_name = $2, father_name = $3 WHERE id = $4`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		teacherNew.FirstName,
		teacherNew.LastName,
		teacherNew.FatherName,
		teacherOld.Id,
	)
	return err
}

func (s *Storage) DeleteTeacher(teacher models.Teacher) error {
	query := `DELETE FROM people WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, teacher.Id)
	return err
}
