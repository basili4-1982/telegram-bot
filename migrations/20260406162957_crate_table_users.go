package migrations

import (
	"context"
	"database/sql"

	goose "github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCrateTableUsers, downCrateTableUsers)
}

func upCrateTableUsers(ctx context.Context, tx *sql.Tx) error {
	const q = `create table rate
(
    id         varchar(50)                                  not null
        constraint rate_pk
            primary key,
    rate       double precision                             not null,
    unit       text                     default 'BTC'::text not null,
    up         smallint                 default 100         not null,
    down       smallint                 default 100         not null,
    created_at timestamp with time zone default now()       not null,
    update_at  timestamp with time zone default now()       not null
);`
	_, err := tx.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func downCrateTableUsers(ctx context.Context, tx *sql.Tx) error {
	const q = `drop table rate;`
	_, err := tx.ExecContext(ctx, q)

	if err != nil {
		return err
	}

	return nil
}
