package swap_adjacent_in_lr_string

import (
	"math/rand"
	"testing"
)

func TestGenRandStartEnd(t *testing.T) {
	n := 10000
	options := []byte{'R', 'X', 'L'}
	start, end := make([]byte, n), make([]byte, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(3)
		start[i] = options[x]
		x = rand.Intn(3)
		end[i] = options[x]
	}
	t.Log(string(start))
	t.Log(string(end))
}

func TestSalrs(t *testing.T) {
	start, end := "XLLR", "LXLX"
	t.Log(canTransform(start, end))
}
