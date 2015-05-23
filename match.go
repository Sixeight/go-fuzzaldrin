package fuzzaldrin

import (
	"math"
	"os"
	"unicode"
)

func basenameMatch(str, query string) []int {
	strRunes := []rune(str)

	index := len(strRunes) - 1
	for index >= 0 && os.IsPathSeparator(uint8(strRunes[index])) {
		index--
	}
	slashCount := 0
	lastCharacter := index
	var base string
	for index >= 0 {
		if os.IsPathSeparator(uint8(strRunes[index])) {
			slashCount++
			if base == "" {
				base = string(strRunes[index+1 : lastCharacter+1])
			}
		} else if index == 0 {
			if base == "" {
				if lastCharacter < len(strRunes)-1 {
					base = string(strRunes[0 : lastCharacter+1])
				} else {
					base = string(strRunes)
				}
			}
		}
		index--
	}

	return match(base, query, len(str)-len(base))
}

func match(str, query string, offset int) (indexes []int) {
	strRunes := []rune(str)

	if str == query {
		for i := offset; i < len(strRunes)+offset; i++ {
			indexes = append(indexes, i)
		}
		return
	}

	indexInStr := 0

	for _, ch := range query {
		lcIndex := findIndex(strRunes, unicode.ToLower(ch))
		ucIndex := findIndex(strRunes, unicode.ToUpper(ch))

		minIndex := int(math.Min(float64(lcIndex), float64(ucIndex)))
		if minIndex == -1 {
			minIndex = int(math.Max(float64(lcIndex), float64(ucIndex)))
		}
		if minIndex == -1 {
			return
		}
		indexes = append(indexes, offset+minIndex)
		indexInStr = minIndex

		offset += indexInStr + 1
		strRunes = strRunes[indexInStr+1 : len(strRunes)]

	}

	return
}
