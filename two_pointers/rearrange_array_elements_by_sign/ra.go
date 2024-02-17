package rearrange_array_elements_by_sign

func rearrangeArray(nums []int) []int {
	r := make([]int, len(nums))

	var pc, nc int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			r[pc<<1] = nums[i]
			pc++
		} else {
			r[nc<<1+1] = nums[i]
			nc++
		}
	}

	return r
}
