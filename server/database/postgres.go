package database

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(ctx context.Context, dsn string) (*sqlx.DB, error) {
	conn, err := sqlx.ConnectContext(ctx, "pgx", dsn)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(40)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxIdleTime(2 * time.Minute)
	return conn, nil
}
