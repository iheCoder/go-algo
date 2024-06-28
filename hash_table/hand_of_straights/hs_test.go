package hand_of_straights

import "testing"

func TestHS(t *testing.T) {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	gs := 3
	t.Log(isNStraightHand(hand, gs))
}
