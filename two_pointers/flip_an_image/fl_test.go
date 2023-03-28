package flip_an_image

import "testing"

func TestInvert(t *testing.T) {
	x := 0
	t.Log(1 ^ x)
	x = 1
	t.Log(1 ^ x)
}

func TestFLIJK(t *testing.T) {
	flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}})
}
