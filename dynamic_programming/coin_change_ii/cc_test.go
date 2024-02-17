package coin_change_ii

import "testing"

func TestCC(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 5
	t.Log(change(amount, coins))
}
