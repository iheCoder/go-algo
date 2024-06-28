package lemonade_change

import "testing"

func TestLC(t *testing.T) {
	bills := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	t.Log(lemonadeChange(bills))
}
