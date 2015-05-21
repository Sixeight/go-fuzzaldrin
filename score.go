package fuzzaldrin

import (
	"math"
	"os"
	"strings"
)

func score(str, query string) float64 {
	if str == query {
		return 1
	}

	if os.IsPathSeparator(str[len(str)-1]) {
		return 1
	}

	totalScore := 0.0
	queryLength := len(query)
	strLength := len(str)

	indexInQuery := 0
	indexInStr := 0

	for indexInQuery < queryLength {
		ch := query[indexInQuery]
		indexInQuery++
		lcIndex := strings.Index(str, strings.ToLower(string(ch)))
		ucIndex := strings.Index(str, strings.ToUpper(string(ch)))
		minIndex := math.Min(float64(lcIndex), float64(ucIndex))
		if minIndex == -1 {
			minIndex = math.Max(float64(lcIndex), float64(ucIndex))
		}
		indexInStr = int(minIndex)
		if indexInStr == -1 {
			return 0
		}

		chScore := 0.1

		if str[indexInStr] == ch {
			chScore += 0.1
		}

		if indexInStr == 0 || os.IsPathSeparator(str[indexInStr-1]) {
			// Start of string bonus
			chScore += 0.8
		} else if strings.Contains("-_ ", string(str[indexInStr-1])) {
			// Start of word bonus
			chScore += 0.7
		}

		// Trim string to after current abbreviation match
		str = str[indexInStr+1 : strLength]
		strLength = len(str)

		totalScore += chScore
	}

	queryScore := totalScore / float64(queryLength)
	return ((queryScore * (float64(queryLength) / float64(strLength))) + queryScore) / 2.0
}
