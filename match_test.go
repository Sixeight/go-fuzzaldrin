package fuzzaldrin

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"
)

func testMatchedIndexes(result, expected []int, t *testing.T) {
	if len(result) == 0 && len(expected) == 0 {
		return
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v should be %v", result, expected)
	}
}

func TestMatch(t *testing.T) {
	testMatchedIndexes(Match("Hello World", "he"), []int{0, 1}, t)
	testMatchedIndexes(Match("", ""), []int{}, t)
	testMatchedIndexes(Match("Hello World", "wor"), []int{6, 7, 8}, t)
	testMatchedIndexes(Match("Hello World", "d"), []int{10}, t)
	testMatchedIndexes(Match("Hello World", "elwor"), []int{1, 2, 6, 7, 8}, t)
	testMatchedIndexes(Match("Hello World", "er"), []int{1, 8}, t)
	testMatchedIndexes(Match("Hello World", ""), []int{}, t)
	testMatchedIndexes(Match("", "abc"), []int{}, t)
}

func TestMatch_Path(t *testing.T) {
	testMatchedIndexes(Match(path.Join("X", "Y"), path.Join("X", "Y")), []int{0, 1, 2}, t)
	testMatchedIndexes(Match(path.Join("X", "X-x"), "X"), []int{0, 2}, t)
	testMatchedIndexes(Match(path.Join("X", "Y"), "XY"), []int{0, 2}, t)
	testMatchedIndexes(Match(path.Join("-", "X"), "X"), []int{2}, t)
	testMatchedIndexes(Match(path.Join("X-", "-"), fmt.Sprintf("X%s", string(os.PathSeparator))), []int{0, 2}, t)
}

func TestMatch_DoubleMatch(t *testing.T) {
	testMatchedIndexes(Match(path.Join("XY", "XY"), "XY"), []int{0, 1, 3, 4}, t)
	testMatchedIndexes(Match(path.Join("--X-Y-", "-X--Y"), "XY"), []int{2, 4, 8, 11}, t)
}
