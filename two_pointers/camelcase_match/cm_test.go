package camelcase_match

import (
	"reflect"
	"testing"
)

type testData struct {
	queries  []string
	pattern  string
	expected []bool
}

func TestCM(t *testing.T) {
	tds := []testData{
		{
			queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern:  "FB",
			expected: []bool{true, false, true, true, false},
		},
	}

	for i, td := range tds {
		r := camelMatch(td.queries, td.pattern)
		if !reflect.DeepEqual(r, td.expected) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
