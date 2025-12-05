package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
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
