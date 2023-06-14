package divisor_game

import "testing"

func TestDg(t *testing.T) {
	n := 12
	t.Log(divisorGame(n))
	n = 8
	t.Log(divisorGame(n))
	n = 6
	t.Log(divisorGame(n))
}
