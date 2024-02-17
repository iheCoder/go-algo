package fruit_into_baskets

import "testing"

func TestTF(t *testing.T) {
	fruits := []int{0, 1, 1, 4, 3}
	t.Log(totalFruit(fruits))
}
