package solve_the_equation

import "testing"

func TestSE(t *testing.T) {
	e := "x+5-3+x=6+x-2"
	t.Log(solveEquation(e))
}
