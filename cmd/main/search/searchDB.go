package search

import (
	"context"
	"gosearch/logger"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Result struct {
	URL   string
	Score float64
}

var (
	ctx = context.Background()
	defaultDatabaseUrl = "postgres://user:pass@host.docker.internal:5433/goprocess"
)

func getDBUrl() string {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = defaultDatabaseUrl
	}

	return databaseUrl
}

func Search(query string) []Result {
	dbUrl := getDBUrl()
	connection, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		logger.Log.Error("Unable to get connection pool", "error", err)
		return nil
	}
	defer connection.Close()

	terms := strings.Fields(strings.ToLower(query))
	if len(terms) == 0 {
		return nil
	}

	rows, err := connection.Query(ctx, `
		WITH matched AS (
			SELECT url, term, termcount
			FROM t_url_term_count
			WHERE term = ANY($1)
		),
		term_stats AS (
			SELECT term, COUNT(DISTINCT url) AS docs_with_term
			FROM matched
			GROUP BY term
		)
		SELECT
			m.url,
			SUM(m.termcount * LN(tm.value::float / ts.docs_with_term)) AS score
		FROM matched m
		JOIN term_stats ts ON ts.term = m.term
		CROSS JOIN t_metadata tm
		WHERE tm.key = 'total_urls'
		GROUP BY m.url
		ORDER BY score DESC
		LIMIT 10;
	`, terms)
	if err != nil {
		logger.Log.Error("Search query failed", "error", err)
		return nil
	}
	defer rows.Close()

	var results []Result
	for rows.Next() {
		var r Result
		if err := rows.Scan(&r.URL, &r.Score); err != nil {
			logger.Log.Error("Failed to scan row", "error", err)
			continue
		}
		results = append(results, r)
	}

	return results
}
