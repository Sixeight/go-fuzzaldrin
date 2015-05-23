package fuzzaldrin

func findIndex(strRunes []rune, ch rune) int {
	for i, c := range strRunes {
		if c == ch {
			return i
		}
	}
	return -1
}

func contains(intArray []int, search int) bool {
	for _, i := range intArray {
		if i == search {
			return true
		}
	}
	return false
}
