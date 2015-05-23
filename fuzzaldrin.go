package fuzzaldrin

import (
	"os"
	"strings"
)

// Filter filters candidates with query
func Filter(candidateds []string, query string, maxResults int) []string {
	queryHasSlashes := false
	if query != "" {
		queryHasSlashes = strings.Index(query, string(os.PathSeparator)) != -1
		query = strings.Replace(query, " ", "", -1)
	}
	return filter(candidateds, query, queryHasSlashes, maxResults)
}

// Score caliculates mathing score
func Score(str, query string) float64 {
	if str == "" {
		return 0
	}
	if query == "" {
		return 0
	}
	if str == query {
		return 2
	}

	queryHasSlashes := strings.Index(query, string(os.PathSeparator)) != -1
	query = strings.Replace(query, " ", "", -1)
	calcScore := score(str, query)
	if !queryHasSlashes {
		calcScore = basenameScore(str, query, calcScore)
	}
	return calcScore
}
