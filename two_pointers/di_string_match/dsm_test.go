package di_string_match

import (
	"testing"
)

type testData struct {
	s        string
	expected []int
}

func TestDSM(t *testing.T) {
	tds := []testData{
		{
			s: "IID",
		},
	}

	for _, td := range tds {
		r := diStringMatchDep(td.s)
		t.Log(r)
	}
}
