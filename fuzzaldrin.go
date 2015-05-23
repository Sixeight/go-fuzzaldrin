package fuzzaldrin

import (
	"os"
	"sort"
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

// Match returns matched indexes
func Match(str, query string) (indexes []int) {
	if str == "" {
		return
	}
	if query == "" {
		return
	}

	strRunes := []rune(str)

	if str == query {
		for i := 0; i < len(strRunes); i++ {
			indexes = append(indexes, i)
		}
		return
	}

	queryHasSlashes := strings.Index(query, string(os.PathSeparator)) != -1
	query = strings.Replace(query, " ", "", -1)
	indexes = match(str, query, 0)
	if !queryHasSlashes {
		baseIndexes := basenameMatch(str, query)
		for _, i := range baseIndexes {
			if contains(indexes, i) {
				continue
			}
			indexes = append(indexes, i)
		}
		sort.Ints(indexes)
	}

	return
}
