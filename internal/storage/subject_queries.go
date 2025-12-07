package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetSubjects() ([]models.Subject, error) {
	query := `SELECT * FROM subjects`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Subject])
}

func (s *Storage) InsertSubject(subject models.Subject) error {
	query := `INSERT INTO subjects (name) VALUES($1)`
	_, err := s.pool.Exec(context.Background(), query, subject.Name)
	return err
}

func (s *Storage) UpdateSubject(subjectOld, subjectNew models.Subject) error {
	query := `UPDATE subjects SET name = $1 WHERE id = $2`
	_, err := s.pool.Exec(context.Background(), query, subjectNew.Name, subjectOld.Id)
	return err
}

func (s *Storage) DeleteSubject(subject models.Subject) error {
	query := `DELETE FROM subjects WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, subject.Id)
	return err
}
