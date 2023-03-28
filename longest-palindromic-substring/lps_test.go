package longest_palindromic_substring

import "testing"

func TestLPSTwoPointers(t *testing.T) {
	type ts struct {
		input    string
		expected string
	}

	td := []ts{
		{
			input:    "babad",
			expected: "aba",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "a",
			expected: "a",
		},
		{
			input:    "aacabdkacaa",
			expected: "aca",
		},
	}

	for _, ti := range td {
		r := longestPalindrome(ti.input)
		if !checkExpected(r, ti.expected) {
			t.Fatalf("expect %s got %s ", ti.expected, r)
		}
	}
}

func checkExpected(output, expected string) bool {
	if output == expected {
		return true
	}

	if len(output) != len(expected) {
		return false
	}

	i := 0
	j := len(output) - 1
	for i <= j {
		if output[i] != output[j] {
			return false
		}
		i++
		j--
	}

	return true
}
