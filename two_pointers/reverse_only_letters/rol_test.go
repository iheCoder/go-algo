package reverse_only_letters

import "testing"

type testData struct {
	s        string
	expected string
}

func TestROL(t *testing.T) {
	tds := []testData{
		{
			s:        "ab-cd",
			expected: "dc-ba",
		},
		{
			s:        "a-bC-dEf-ghIj",
			expected: "j-Ih-gfE-dCba",
		},
		{
			s:        "Test1ng-Leet=code-Q!",
			expected: "Qedo1ct-eeLg=ntse-T!",
		},
		{
			s:        "----------",
			expected: "----------",
		},
		{
			s:        "abcde",
			expected: "edcba",
		},
		{
			s:        "---a",
			expected: "---a",
		},
	}

	for i, td := range tds {
		r := reverseOnlyLetters(td.s)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
