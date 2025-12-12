package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetStudents() ([]models.Student, error) {
	query := `SELECT * FROM students_with_group`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Student])
}

func (s *Storage) InsertStudent(student models.Student) error {
	query := `insert into students_with_group (first_name, last_name, father_name, name) values($1, $2, $3, $4);`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		student.FirstName,
		student.LastName,
		student.FatherName,
		student.Group,
	)
	return err
}

func (s *Storage) UpdateStudent(student, studentNew models.Student) error {
	query := `UPDATE students_with_group SET first_name = $1, last_name = $2, father_name = $3, name = $4 WHERE id = $5`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		studentNew.FirstName,
		studentNew.LastName,
		studentNew.FatherName,
		studentNew.Group,
		student.Id,
	)
	return err
}

func (s *Storage) DeleteStudent(student models.Student) error {
	query := `DELETE FROM people WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, student.Id)
	return err
}

func (s *Storage) GetStudentsNoYearGroup() ([]models.Student, error) {
	query := `SELECT * FROM students_with_group_no_year`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var first, last, father, group string
	students := make([]models.Student, 0)
	pgx.ForEachRow(rows, []any{&first, &last, &father, &group}, func() error {
		students = append(students, models.Student{-1, first, last, father, group})
		return nil
	})
	return students, nil
}
