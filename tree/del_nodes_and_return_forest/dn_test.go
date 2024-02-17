package del_nodes_and_return_forest

import (
	"math"
	"testing"
)

func TestArrToMap(t *testing.T) {
	to_delete := []int{3, 5}
	dm := make(map[int]bool)
	for _, td := range to_delete {
		dm[td] = true
	}
	t.Log(dm)
	t.Log(math.Pow10(9) + 7)
}
