package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(connection string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &Storage{pool}, nil
}

func (s *Storage) HashUser(login, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`

	_, err = s.pool.Query(context.Background(), query, login, string(hash))
	if err != nil {
		return err
	}
	return nil
}
