package fuzzaldrin

import "testing"

func TestScore(t *testing.T) {
	first := Score("Hello World", "he")
	second := Score("Hello World", "Hello")
	if first >= second {
		t.Errorf("%v is greater than %v", first, second)
	}

	justMatch := Score("Hello World", "Hello World")
	if justMatch != 2 {
		t.Errorf("%v must be 2", justMatch)
	}

	emptyQuery := Score("Hello World", "")
	if emptyQuery != 0 {
		t.Errorf("%v must be 0", emptyQuery)
	}

	emptyString := Score("", "he")
	if emptyString != 0 {
		t.Errorf("%v must be 0", emptyString)
	}

	japaneseString1 := Score("こんにちは World", "こん")
	japaneseString2 := Score("こんにちは World", "にちは")
	if japaneseString1 <= japaneseString2 {
		t.Errorf("%v must be grater than %v", japaneseString1, japaneseString2)
	}

	japaneseString3 := Score("こんにちは World", "World")
	if japaneseString3 <= japaneseString2 {
		t.Errorf("%v must be grater than %v", japaneseString3, japaneseString2)
	}

	spaceAsSeparate := Score("bar/foo", "br f")
	if spaceAsSeparate == 0 {
		t.Errorf("%v must not be 0", spaceAsSeparate)
	}
}
