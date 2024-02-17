package single_number3

import (
	"strconv"
	"strings"
	"testing"
)

func TestSN3(t *testing.T) {
	t.Log(1<<31 - 1)
	arr := []int{1, 2, 1, 3, 2, 5}
	t.Log(singleNumber(arr))
	var nums []string
	for i := 0; i < 16; i++ {
		nums = append(nums, strconv.Itoa(1<<(i<<1)))
	}
	t.Log(strings.Join(nums, ":true,"))
}
