package add_two_num_ii

import (
	"strconv"
	"testing"
)

func TestAddNumStr(t *testing.T) {
	n1 := "99"
	n2 := "9"
	t.Log(strconv.Itoa('0'))
	t.Log(strconv.Itoa(int(byte(n1[0]))))
	t.Log(addNums(n1, n2))
}
