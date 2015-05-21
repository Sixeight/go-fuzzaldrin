package fuzzaldrin

import (
	"os"
	"strings"
)

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
	str = strings.Replace(str, " ", "", -1)
	calcScore := score(str, query)
	if queryHasSlashes {
		calcScore = basenameScore(str, query, calcScore)
	}
	return calcScore
}
