package db

import (
	"context"

	"github.com/chadsmith12/pacer/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)


func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	dbConfig, err := config.LoadDatabase()
	if err != nil {
		return nil, err
	}

	return ConnectWithConfig(ctx, dbConfig)
}

func ConnectWithConfig(ctx context.Context, dbConfig config.DatabaseConfig) (*pgxpool.Pool, error) {
	pgConfig, err := pgxpool.ParseConfig(dbConfig.String())
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(ctx, pgConfig)
}
