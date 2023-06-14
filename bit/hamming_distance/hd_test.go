package hamming_distance

import "testing"

func TestHD(t *testing.T) {
	x := 1
	y := 4
	t.Log(hammingDistance(x, y))
}
