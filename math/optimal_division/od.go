package optimal_division

import (
	"fmt"
	"strconv"
)

// 4/3/2 最小是4/3/2
// 4/2/3 最小是 4/2/3
// 2/3/4 最小是2/3/4
func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	s := fmt.Sprintf("%d/(", nums[0])
	for i := 1; i < n; i++ {
		s += strconv.Itoa(nums[i]) + "/"
	}
	s = s[:len(s)-1]
	s += ")"
	return s
}
