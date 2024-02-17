package sort_char_by_frequency

import (
	"math/rand"
	"testing"
)

func TestSC(t *testing.T) {
	s := "tree"
	t.Log(frequencySort(s))
	t.Log(byte('a'))
	t.Log(byte('A'))
}

func TestGenRandomChar(t *testing.T) {
	n := 500000
	bs := make([]byte, 0, n)
	var c byte
	for i := 0; i < n; i++ {
		x := rand.Intn(52)
		if x >= 26 {
			c = byte('A' + x - 26)
		} else {
			c = byte('a' + x)
		}
		bs = append(bs, c)
	}
	t.Log(string(bs))
}
