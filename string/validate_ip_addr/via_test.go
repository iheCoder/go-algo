package validate_ip_addr

import (
	"strconv"
	"testing"
)

func TestStrConv(t *testing.T) {
	s := "000"
	num, err := strconv.Atoi(s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}

func TestVIA(t *testing.T) {
	s := "2001:0db8:85a3:00000:0:8A2E:0370:7334"
	t.Log(validIPAddress(s))
}
