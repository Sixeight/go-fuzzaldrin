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

	lastSeparator := Score("/path/to/foo", "foo")
	if lastSeparator != 1 {
		t.Errorf("%v mus be 1", lastSeparator)
	}
}
