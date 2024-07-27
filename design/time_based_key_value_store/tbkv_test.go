package time_based_key_value_store

import "testing"

func TestTM(t *testing.T) {
	tm := Constructor()
	//tm.Set("foo", "bar", 1)
	//t.Log(tm.Get("foo", 3))
	//
	//tm.Set("foo", "bar2", 4)
	//t.Log(tm.Get("foo", 4))

	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	t.Log(tm.Get("love", 15))
}
