package nim_game

import "testing"

func TestNimGame(t *testing.T) {
	n := 10086
	t.Log(canWinNim(n))
	n = 2147483647
	t.Log(canWinNim(n))
	n = 5
	t.Log(canWinNim(n))
}
