package open_the_lock

import "testing"

func TestOpenLock(t *testing.T) {
	x := '9'
	y := '0'
	t.Log(int(x - y))
	t.Log(int(x) - int(y))
	ds := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	t.Log(openLock(ds, target))
}
