package search

import "context"

type Searcher interface {
	Search(query string, ctx context.Context) []Result
}
