package db

import (
	"context"
	"gosearch/logger"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	defaultDatabaseUrl = "postgres://user:pass@host.docker.internal:5433/goprocess"
)

func getDBUrl() string {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = defaultDatabaseUrl
	}

	return databaseUrl
}

func GetDB(ctx context.Context) *pgxpool.Pool {
	dbUrl := getDBUrl()
	connection, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		logger.Log.Error("Unable to get DB connection", "error", err)
		panic("Unaable to get DB connection")
	}

	return connection
}
