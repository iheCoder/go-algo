package bulb_switcher

import (
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	n := 32
	x := sort.Search(n, func(i int) bool {
		return i*i >= n
	})
	t.Log(x)
}
