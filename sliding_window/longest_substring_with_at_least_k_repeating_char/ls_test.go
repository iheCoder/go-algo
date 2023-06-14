package longest_substring_with_at_least_k_repeating_char

import "testing"

func TestLS(t *testing.T) {
	type testData struct {
		s        string
		k        int
		expected int
	}
	tds := []testData{
		//{
		//	s:        "aaabb",
		//	k:        3,
		//	expected: 3,
		//},
		//{
		//	s:        "bbaaacbd",
		//	k:        3,
		//	expected: 3,
		//},
		{
			s:        "a",
			k:        1,
			expected: 1,
		},
	}
	for i, td := range tds {
		got := longestSubstring(td.s, td.k)
		if got != td.expected {
			t.Fatalf("%d s %s k %d exp %d got %d", i, td.s, td.k, td.expected, got)
		}
	}
}
