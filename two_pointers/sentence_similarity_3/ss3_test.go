package sentence_similarity_3

import "testing"

type testData struct {
	s1       string
	s2       string
	expected bool
}

func TestLRSub(t *testing.T) {
	tds := []testData{
		{
			s1:       "My name is Haley",
			s2:       "My Haley",
			expected: true,
		},
	}

	for _, td := range tds {
		r := areSentencesSimilarLRSub(td.s1, td.s2)
		if r != td.expected {
			t.Fatalf("expect %v got %v", td.expected, r)
		}
	}
}

func TestTryCompose(t *testing.T) {
	tds := []testData{
		{
			s1:       "My name is Haley",
			s2:       "My Haley",
			expected: true,
		},
		{
			s1:       "of",
			s2:       "A lot of words",
			expected: false,
		},
		{
			s1:       "Are you Okay",
			s2:       "are you okay",
			expected: false,
		},
		{
			s1:       "B",
			s2:       "ByI BMyQIqce b bARkkMaABi vlR RLHhqjNzCN oXvyK zRXR q ff B yHS OD KkvJA P JdWksnH",
			expected: false,
		},
		{
			s1:       "a A",
			s2:       "aAAa",
			expected: false,
		},
		{
			s1:       "A B C D B B",
			s2:       "A B B",
			expected: true,
		},
	}

	for i, td := range tds {
		r := areSentencesSimilarCompose(td.s1, td.s2)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
