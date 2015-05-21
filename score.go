package fuzzaldrin

import (
	"math"
	"os"
	"strings"
)

func basenameScore(str, query string, calcScore float64) float64 {
	index := len(str) - 1
	for os.IsPathSeparator(str[index]) {
		index--
	}
	slashCount := 0
	lastCharacter := index
	var base string
	for index >= 0 {
		if os.IsPathSeparator(str[index]) {
			slashCount++
			base = str[index+1 : lastCharacter+1]
		} else if index == 0 {
			if lastCharacter < len(str)-1 {
				base = str[0 : lastCharacter+1]
			} else {
				base = str
			}
		}
		index--
	}

	// Basename matches count for more.
	if base == str {
		calcScore *= 2
	} else if base != "" {
		calcScore += score(base, query)
	}

	// Shallow files are scored higher
	segmentCount := slashCount + 1
	depth := math.Max(1.0, float64(10-segmentCount))
	calcScore *= depth * 0.01
	return calcScore
}

func score(str, query string) float64 {
	if str == query {
		return 1
	}

	if queryIsLastPathSegment(str, query) {
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
		str = str[indexInStr+1 : len(str)]

		totalScore += chScore
	}

	queryScore := totalScore / float64(queryLength)
	return ((queryScore * (float64(queryLength) / float64(strLength))) + queryScore) / 2.0
}

func queryIsLastPathSegment(str, query string) bool {
	if len(str) <= len(query) {
		return false
	}
	if os.IsPathSeparator(str[len(str)-len(query)-1]) {
		return strings.Index(str, query) == (len(str) - len(query))
	}
	return false
}
