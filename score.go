package fuzzaldrin

import (
	"math"
	"os"
	"strings"
	"unicode"
)

func basenameScore(str, query string, calcScore float64) float64 {
	index := len(str) - 1
	for index >= 0 && os.IsPathSeparator(str[index]) {
		index--
	}
	slashCount := 0
	lastCharacter := index
	var base string
	for index >= 0 {
		if os.IsPathSeparator(str[index]) {
			slashCount++
			if base == "" {
				base = str[index+1 : lastCharacter+1]
			}
		} else if index == 0 {
			if base == "" {
				if lastCharacter < len(str)-1 {
					base = str[0 : lastCharacter+1]
				} else {
					base = str
				}
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

	strRunes := []rune(str)

	totalScore := 0.0
	queryLength := len([]rune(query))
	strLength := len(strRunes)

	indexInStr := 0

	for _, ch := range query {
		lcIndex := findIndex(strRunes, unicode.ToLower(ch))
		ucIndex := findIndex(strRunes, unicode.ToUpper(ch))
		minIndex := math.Min(float64(lcIndex), float64(ucIndex))
		if minIndex == -1 {
			minIndex = math.Max(float64(lcIndex), float64(ucIndex))
		}
		if minIndex == -1 {
			return 0
		}

		chScore := 0.1

		indexInStr = int(minIndex)
		if strRunes[indexInStr] == ch {
			chScore += 0.1
		}

		if indexInStr == 0 || os.IsPathSeparator(uint8(strRunes[indexInStr-1])) {
			// Start of string bonus
			chScore += 0.8
		} else if strings.Contains("-_ ", string(strRunes[indexInStr-1])) {
			// Start of word bonus
			chScore += 0.7
		}

		// Trim string to after current abbreviation match
		strRunes = strRunes[indexInStr+1 : len(strRunes)]

		totalScore += chScore
	}

	queryScore := totalScore / float64(queryLength)
	return ((queryScore * (float64(queryLength) / float64(strLength))) + queryScore) / 2.0
}

func findIndex(str []rune, ch rune) int {
	for i, c := range str {
		if c == ch {
			return i
		}
	}
	return -1
}

func queryIsLastPathSegment(str, query string) bool {
	if len(str) <= len(query) {
		return false
	}
	if os.IsPathSeparator(str[len(str)-len(query)-1]) {
		return strings.LastIndex(str, query) == (len(str) - len(query))
	}
	return false
}
