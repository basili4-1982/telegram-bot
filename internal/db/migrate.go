package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	_ "telegram-bot/migrations"
)

func Up(db *pgxpool.Pool) error {
	conn := stdlib.GetPoolConnector(db)
	err := goose.Up(sql.OpenDB(conn), "migrations")
	if err != nil {
		return err
	}

	return nil
}
