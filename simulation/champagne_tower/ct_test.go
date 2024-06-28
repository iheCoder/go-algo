package champagne_tower

import "testing"

func TestCT(t *testing.T) {
	p := 7
	r, c := 3, 2
	t.Log(champagneTower(p, r, c))
}
