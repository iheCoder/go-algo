package single_number2

import "testing"

func TestSN(t *testing.T) {
	arr := []int{1, 2, 2, 2}
	t.Log(singleNumber(arr))
}
