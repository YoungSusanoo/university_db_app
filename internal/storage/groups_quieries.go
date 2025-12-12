package storage

import (
	"context"
	"university_app/internal/models"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetGroups() ([]models.Group, error) {
	query := `SELECT id, name FROM groups`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Group])
}

func (s *Storage) InsertGroup(groups models.Group) error {
	query := `INSERT INTO groups (name) VALUES($1)`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		groups.Name,
	)
	return err
}

func (s *Storage) UpdateGroup(group, groupNew models.Group) error {
	query := `UPDATE groups SET name = $1 WHERE id = $2`
	_, err := s.pool.Exec(
		context.Background(),
		query,
		groupNew.Name,
		group.Id,
	)
	return err
}

func (s *Storage) DeleteGroup(group models.Group) error {
	query := `DELETE FROM groups WHERE id = $1`
	_, err := s.pool.Exec(context.Background(), query, group.Id)
	return err
}

func (s *Storage) GetGroupsNoYear() ([]models.Group, error) {
	query := `SELECT DISTINCT split_part(name, '_', 1) FROM groups`
	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var name string
	groups := make([]models.Group, 0)
	pgx.ForEachRow(rows, []any{&name}, func() error {
		groups = append(groups, models.Group{-1, name})
		return nil
	})
	return groups, err
}
