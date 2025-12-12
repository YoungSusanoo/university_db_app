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
		&stud.Group,
		&subj.Name,
		&value,
	},
		func() error {
			marks = append(marks, models.Mark{id, teach, stud, subj, value})
			return nil
		},
	)

	return marks, err
}

func (s *Storage) InsertMark(mark models.Mark) error {
	query := `insert into marks_with_names (
		t_first_name, 
		t_last_name,
		t_father_name,
		s_first_name,
		s_last_name,
		s_father_name,
		group_name,
		subj_name,
		value)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		mark.Teach.FirstName,
		mark.Teach.LastName,
		mark.Teach.FatherName,
		mark.Stud.FirstName,
		mark.Stud.LastName,
		mark.Stud.FatherName,
		mark.Stud.Group,
		mark.Subj.Name,
		mark.Value,
	)
	return err
}

func (s *Storage) UpdateMark(mark, markNew models.Mark) error {
	query := `UPDATE marks_with_names SET 
		t_first_name = $1, 
		t_last_name = $2, 
		t_father_name = $3,
		s_first_name = $4,
		s_last_name = $5,
		s_father_name = $6,
		group_name = $7,
		subj_name = $8,
		value = $9 
		WHERE id = $10`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		markNew.Teach.FirstName,
		markNew.Teach.LastName,
		markNew.Teach.FatherName,
		markNew.Stud.FirstName,
		markNew.Stud.LastName,
		markNew.Stud.FatherName,
		markNew.Stud.Group,
		markNew.Subj.Name,
		markNew.Value,
		mark.Id,
	)
	return err
}

func (s *Storage) DeleteMark(mark models.Mark) error {
	query := `DELETE FROM marks WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, mark.Id)
	return err
}

func (s *Storage) GetAvgGroupRange(start, end int, name string) (avgs []models.YearAverage, err error) {
	query := `SELECT * FROM get_groups_avg_year_range($1, $2, $3)`
	rows, err := s.pool.Query(context.Background(), query, start, end, name)
	if err != nil {
		return nil, err
	}
	avgs = make([]models.YearAverage, 0)
	var year int64
	var avg float32
	pgx.ForEachRow(rows, []any{&year, &avg}, func() error {
		avgs = append(avgs, models.YearAverage{year, avg})
		return nil
	})
	return
}

func (s *Storage) GetAvgStudentRange(start, end int, student models.Student) (avgs []models.YearAverage, err error) {
	query := `SELECT * FROM get_students_avg_year_range($1, $2, $3, $4, $5, $6)`
	rows, err := s.pool.Query(
		context.Background(),
		query,
		start,
		end,
		student.FirstName,
		student.LastName,
		student.FatherName,
		student.Group,
	)
	if err != nil {
		return nil, err
	}
	avgs = make([]models.YearAverage, 0)
	var year int64
	var avg float32
	pgx.ForEachRow(rows, []any{&year, &avg}, func() error {
		avgs = append(avgs, models.YearAverage{year, avg})
		return nil
	})
	return
}

func (s *Storage) GetAvgTeacherRange(start, end int, teacher models.Teacher) (avgs []models.YearAverage, err error) {
	query := `SELECT * FROM get_teacher_avg_year_range($1, $2, $3, $4, $5)`
	rows, err := s.pool.Query(
		context.Background(),
		query,
		start,
		end,
		teacher.FirstName,
		teacher.LastName,
		teacher.FatherName,
	)
	if err != nil {
		return nil, err
	}
	avgs = make([]models.YearAverage, 0)
	var year int64
	var avg float32
	pgx.ForEachRow(rows, []any{&year, &avg}, func() error {
		avgs = append(avgs, models.YearAverage{year, avg})
		return nil
	})
	return
}
