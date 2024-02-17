package restore_ip_addr

import "testing"

func TestRia(t *testing.T) {
	s := "25525511135"
	//t.Log(restoreIpAddresses(s))
	s = "101023"
	//t.Log(restoreIpAddresses(s))
	s = "5561122202"
	t.Log(restoreIpAddresses(s))
}
