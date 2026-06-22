package search

import (
	"context"
	"gosearch/logger"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBSearcher struct {
	Connection *pgxpool.Pool
}

type Result struct {
	URL   string
	Score float64
}

func (dbSearcher DBSearcher) Search(query string, ctx context.Context) []Result {
	if dbSearcher.Connection == nil {
		panic("Null connection specified")
	}

	terms := strings.Fields(strings.ToLower(query))
	if len(terms) == 0 {
		return nil
	}

	rows, err := dbSearcher.Connection.Query(ctx, `
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
