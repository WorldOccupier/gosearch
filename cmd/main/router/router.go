package router

import (
	"context"
	"gosearch/logger"
	"gosearch/search"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	portString = ":8080"
	searchString = "search"
	query = "query"
)

func searchQuery(searcher search.Searcher, ctx context.Context) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		queryString := ginContext.Query(query)
		results := searcher.Search(queryString, ctx)
		ginContext.JSON(http.StatusOK, results)
	}
}

func Run(searcher search.Searcher, ctx context.Context) {
	router := gin.Default()
	router.GET(searchString, searchQuery(searcher, ctx))
	if err := router.Run(portString); err != nil {
		logger.Log.Error("Server failed to start", "error", err)
	}
}
