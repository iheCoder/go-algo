package evaluate_division

import "testing"

func TestED(t *testing.T) {
	es := [][]string{{"a", "b"}}
	values := []float64{0.5}
	qs := [][]string{
		{
			"a",
			"b",
		},
	}
	t.Log(calcEquation(es, values, qs))
}
