package can_i_win

import "testing"

func TestCanIWin(t *testing.T) {
	mci := 10
	total := 11
	t.Log(canIWin(mci, total))
}
