package broken_calculator

import "testing"

func TestBC(t *testing.T) {
	sv, target := 2, 93809375
	t.Log(brokenCalc(sv, target))
}
