package expressive_words

import "testing"

type testData struct {
	s        string
	words    []string
	expected int
}

func TestEW(t *testing.T) {
	tds := []testData{
		//{
		//	s:        "heeellooo",
		//	words:    []string{"hello", "hi", "helo"},
		//	expected: 1,
		//},
		//{
		//	s:        "hellllo",
		//	words:    []string{"hello"},
		//	expected: 1,
		//},
		//{
		//	s:        "helllllo",
		//	words:    []string{"hello"},
		//	expected: 1,
		//},
		//{
		//	s:        "helllloooooox",
		//	words:    []string{"hello"},
		//	expected: 0,
		//},
		//{
		//	s:        "hhhhhhhell",
		//	words:    []string{"hello"},
		//	expected: 0,
		//},
		//{
		//	s:        "aaa",
		//	words:    []string{"aaaa"},
		//	expected: 0,
		//},
		//{
		//	s:        "aaaa",
		//	words:    []string{"aaa"},
		//	expected: 1,
		//},
		//{
		//	s:     "heeellooo",
		//	words: []string{"heeelloooworld"},
		//	expected: 0,
		//},
		{
			s:        "vvkkkejjjjjj",
			words:    []string{"vvkeej", "vkeejj", "vkeej", "vvkejj", "vkej", "vvkkeej"},
			expected: 1,
		},
		{
			s:        "x",
			words:    []string{"x", "xxx", "xx"},
			expected: 1,
		},
	}

	for i, td := range tds {
		r := expressiveWords(td.s, td.words)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
