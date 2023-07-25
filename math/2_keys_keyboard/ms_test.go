package __keys_keyboard

import "testing"

type testData struct {
	n   int
	exp int
}

func TestMS(t *testing.T) {
	tds := []testData{
		{
			n:   8,
			exp: 6,
		},
		{
			n:   9,
			exp: 6,
		},
	}
	for _, td := range tds {
		got := minSteps(td.n)
		if got != td.exp {
			t.Fatalf("n %d got %d exp %d", td.n, got, td.exp)
		}
	}
}
