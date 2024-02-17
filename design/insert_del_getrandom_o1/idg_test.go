package insert_del_getrandom_o1

import "testing"

func TestIDG(t *testing.T) {
	rs := Constructor()
	rsp := &rs
	rsp.Insert(0)
	rsp.Insert(1)
	rsp.Remove(0)
	rsp.Insert(2)
	rsp.Remove(1)
	t.Log(rsp.GetRandom())
}
