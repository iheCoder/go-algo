package largest_merge_of_two_strings

import (
	"reflect"
	"testing"
)

type testData struct {
	word1    string
	word2    string
	expected string
}

func TestLM4P(t *testing.T) {
	tds := []testData{
		{
			"abcabc",
			"abdcaba",
			"abdcabcabcaba",
		},
		{
			"cabaa",
			"bcaaa",
			"cbcabaaaaa",
		},
		{
			"a",
			"a",
			"aa",
		},
		{
			"bcd",
			"bcdbcddaa",
			"bcdbcddbcdaa",
		},
		{
			"uuurruuuruuuuuuuuruuuuu",
			"urrrurrrrrrrruurrrurrrurrrrruu",
			"uuuurruuuruuuuuuuuruuuuurrrurrrrrrrruurrrurrrurrrrruu",
		},
	}

	for i, td := range tds {
		r := largestMerge4P(td.word1, td.word2)
		if !reflect.DeepEqual(td.expected, r) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestComputeCharClasses(t *testing.T) {
	s := "hello"
	//order := sortChars(s)
	//t.Log(order)
	//c := computeCharClasses(s, order)
	//t.Log(c)
	//d := sortDoubled(s, 1, order, c)
	//t.Log(d)
	sa := buildSuffixArray(s)
	t.Log(sa)
}
