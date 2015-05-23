package fuzzaldrin

import (
	"fmt"
	"os"
	p "path"
	"reflect"
	"runtime"
	"testing"
)

var sep = string(os.PathSeparator)

func buildPath(segment ...string) (path string) {
	if runtime.GOOS == "windows" {
		path = "C:\\\\"
	} else {
		path = "/"
	}
	for _, s := range segment {
		if os.IsPathSeparator(s[0]) {
			path += s
		} else {
			path = p.Join(path, s)
		}
	}
	return
}

func testBestMatch(candidates []string, query, expected string, t *testing.T) {
	results := Filter(candidates, query, 1)[0]
	if results != expected {
		t.Errorf("%v should be %v", results, expected)
	}
}

func TestFilter_General(t *testing.T) {
	candidates := []string{"Gruntfile", "filter", "bile", ""}
	results := Filter(candidates, "file", -1)
	expects := []string{"filter", "Gruntfile"}
	if !reflect.DeepEqual(results, expects) {
		t.Errorf("%v should be %v", results, expects)
	}

	results = Filter(candidates, "file", 1)
	expects = []string{"filter"}
	if !reflect.DeepEqual(results, expects) {
		t.Errorf("%v should be %v", results, expects)
	}
}

func TestFilter_Path(t *testing.T) {
	candidates := []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
	}
	testBestMatch(candidates, "bar", candidates[1], t)

	candidates = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar", sep, sep, sep, sep, sep),
	}
	testBestMatch(candidates, "bar", candidates[1], t)

	candidates = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
		"bar",
	}
	testBestMatch(candidates, "bar", candidates[2], t)

	candidates = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
		buildPath("bar"),
	}
	testBestMatch(candidates, "bar", candidates[2], t)

	candidates = []string{
		buildPath("bar", "foo"),
		fmt.Sprintf("bar%s%s%s%s%s", sep, sep, sep, sep, sep),
	}
	testBestMatch(candidates, "bar", candidates[1], t)

	candidates = []string{
		p.Join("f", "o", "1_a_z"),
		p.Join("f", "o", "a_z"),
	}
	testBestMatch(candidates, "az", candidates[1], t)

	candidates = []string{
		p.Join("f", "1_a_z"),
		p.Join("f", "o", "a_z"),
	}
	testBestMatch(candidates, "az", candidates[1], t)
}

func TestFilter_OnlySep(t *testing.T) {
	candidates := []string{sep}
	results := Filter(candidates, "bar", 1)
	expect := []string{}
	if reflect.DeepEqual(results, expect) {
		t.Errorf("%v should be %v", results, expect)
	}
}

func TestFilter_WithSpace(t *testing.T) {
	candidates := []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
	}
	testBestMatch(candidates, "br f", candidates[0], t)

	candidates = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar foo"),
	}
	testBestMatch(candidates, "br f", candidates[1], t)

	candidates = []string{
		buildPath("barfoo", "foo"),
		buildPath("foo", "barfoo"),
	}
	testBestMatch(candidates, "br f", candidates[1], t)

	candidates = []string{
		buildPath("lib", "exportable.rb"),
		buildPath("app", "models", "table.rb"),
	}
	testBestMatch(candidates, "table", candidates[1], t)
}

func TestFilter_MixedCase(t *testing.T) {
	candidates := []string{"statusurl", "StatusUrl"}

	testBestMatch(candidates, "Status", "StatusUrl", t)
	testBestMatch(candidates, "SU", "StatusUrl", t)
	testBestMatch(candidates, "status", "statusurl", t)
	testBestMatch(candidates, "su", "statusurl", t)
	testBestMatch(candidates, "statusurl", "statusurl", t)
	testBestMatch(candidates, "StatusUrl", "StatusUrl", t)
}

func TestFilter_WithSign(t *testing.T) {
	candidates := []string{"sub-zero", "sub zero", "sub_zero"}
	testBestMatch(candidates, "sz", candidates[0], t)

	candidates = []string{"sub zero", "sub_zero", "sub-zero"}
	testBestMatch(candidates, "sz", candidates[0], t)

	candidates = []string{"sub_zero", "sub-zero", "sub zero"}
	testBestMatch(candidates, "sz", candidates[0], t)

	candidates = []string{"a_b_c", "a_b"}
	testBestMatch(candidates, "ab", candidates[1], t)

	candidates = []string{"z_a_b", "a_b"}
	testBestMatch(candidates, "ab", candidates[1], t)

	candidates = []string{"a_b_c", "c_a_b"}
	testBestMatch(candidates, "ab", candidates[0], t)
}

func TestFilter_DirectoryDepth(t *testing.T) {
	candidates := []string{
		buildPath("app", "models", "sutomotive", "car.rb"),
		buildPath("spec", "factories", "cars.rb"),
	}
	testBestMatch(candidates, "car.rb", candidates[0], t)

	candidates = []string{
		buildPath("app", "models", "sutomotive", "car.rb"),
		"car.rb",
	}
	testBestMatch(candidates, "car.rb", candidates[1], t)

	candidates = []string{
		"car.rb",
		buildPath("app", "models", "sutomotive", "car.rb"),
	}
	testBestMatch(candidates, "car.rb", candidates[0], t)

	candidates = []string{
		buildPath("app", "models", "cars", "car.rb"),
		buildPath("spec", "cars.rb"),
	}
	testBestMatch(candidates, "car.rb", candidates[0], t)
}
