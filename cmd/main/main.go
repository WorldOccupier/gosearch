package main

import (
	"context"
	"gosearch/db"
	"gosearch/router"
	"gosearch/search"
)

var (
	ctx = context.Background()
)

func getSearcher() search.Searcher {
	return search.DBSearcher {
		Connection: db.GetDB(ctx),
	}
}

func main() {
	var searcher = getSearcher()
	router.Run(searcher, ctx)
}
