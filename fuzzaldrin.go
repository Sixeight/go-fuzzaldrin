package fuzzaldrin

import "strings"

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

	str = strings.Replace(str, " ", "", -1)
	calcScore := score(str, query)
	return calcScore
}
