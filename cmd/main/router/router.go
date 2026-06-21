package router

import (
	"gosearch/search"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	portString = ":8080"
	searchString = "search"
	query = "query"
)

func searchQuery(ginContext *gin.Context) {
	queryString := ginContext.Query(query)
	results := search.Search(queryString)
	ginContext.JSON(http.StatusOK, results)
}

func Run() {
	router := gin.Default()
	router.GET(searchString, searchQuery)
	router.Run(portString)
}
