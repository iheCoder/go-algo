package reveal_cards_in_increasing_order

import "testing"

func TestRC(t *testing.T) {
	arr := []int{17, 13, 11, 2, 3, 5, 7}
	t.Log(deckRevealedIncreasing(arr))
}
