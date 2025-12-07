package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetStudents() ([]models.Student, error) {
	query := `SELECT people.id, first_name, last_name, father_name, g.name FROM people, groups g WHERE type = 'S' AND g.id = group_id`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Student])
}

func (s *Storage) InsertStudent(students models.Student) error {
	query := `insert into students_with_group (first_name, last_name, father_name, name) values($1, $2, $3, $4);`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		students.FirstName,
		students.LastName,
		students.FatherName,
		students.Group,
	)
	return err
}

func (s *Storage) UpdateStudent(student, studentNew models.Student) error {
	query := `UPDATE people SET first_name = $1, last_name = $2, father_name = $3 WHERE id = $4`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		studentNew.FirstName,
		studentNew.LastName,
		studentNew.FatherName,
		studentNew.Id,
	)
	return err
}

func (s *Storage) DeleteStudent(student models.Student) error {
	query := `DELETE FROM people WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, student.Id)
	return err
}
