package reverse_words_in_a_string

import (
	"fmt"
	"strings"
	"testing"
)

type testData struct {
	s        string
	expected string
}

func TestRWSCompose(t *testing.T) {
	tds := []testData{
		{
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
		{
			s:        "     s     ",
			expected: "s",
		},
	}

	for i, td := range tds {
		r := reverseWords(td.s)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestStrSplit(t *testing.T) {
	s := "hello  world"
	sr := strings.Split(s, " ")
	fmt.Println(sr)
}

func TestStringMut(t *testing.T) {
	//s:="hello"
	//s[0]='x'
}
