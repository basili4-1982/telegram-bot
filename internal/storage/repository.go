package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{
		db: db,
	}
}

func (s Storage) AddRate(ctx context.Context, userID string, rate float64) error {
	const q = `INSERT INTO rate (id, rate) VALUES ($1, $2)`

	_, err := s.db.Exec(ctx, q, userID, rate)
	if err != nil {
		return fmt.Errorf("add rate: %w", err)
	}

	return nil
}
