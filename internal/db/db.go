package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"telegram-bot/internal/config"
)

func OpenDb(ctx context.Context, cfg *config.DB) (*pgxpool.Pool, error) {
	if cfg == nil {
		panic("db config is nil")
	}

	return pgxpool.New(ctx, cfg.DSN)
}
