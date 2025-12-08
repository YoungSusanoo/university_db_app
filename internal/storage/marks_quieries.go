package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetMarks() ([]models.Mark, error) {
	query := `SELECT * from marks_with_names`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var teach models.Teacher
	var stud models.Student
	var subj models.Subject
	var id int64
	var value int
	marks := make([]models.Mark, 0)
	_, err = pgx.ForEachRow(rows, []any{
		&id,
		&teach.FirstName,
		&teach.LastName,
		&teach.FatherName,
		&stud.FirstName,
		&stud.LastName,
		&stud.FatherName,
		&subj.Name,
		&value,
	},
		func() error {
			marks = append(marks, models.Mark{id, stud, teach, subj, value})
			return nil
		},
	)

	return marks, err
}

// func (s *Storage) InsertMark(marks models.Mark) error {
// 	query := `insert into marks_with_group (first_name, last_name, father_name, name) values($1, $2, $3, $4);`
// 	_, err := s.pool.Exec(
// 		context.Background(),
// 		query,
// 		marks.FirstName,
// 		marks.LastName,
// 		marks.FatherName,
// 		marks.Group,
// 	)
// 	return err
// }

// func (s *Storage) UpdateMark(mark, markNew models.Mark) error {
// 	query := `UPDATE people SET first_name = $1, last_name = $2, father_name = $3 WHERE id = $4`
// 	_, err := s.pool.Exec(
// 		context.Background(),
// 		query,
// 		markNew.FirstName,
// 		markNew.LastName,
// 		markNew.FatherName,
// 		markNew.Id,
// 	)
// 	return err
// }

func (s *Storage) DeleteMark(mark models.Mark) error {
	query := `DELETE FROM marks WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, mark.Id)
	return err
}
