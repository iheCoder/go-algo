package product_of_array_except_self

// 如果求总积，那么存在一个问题就是如果当前元素是0怎么办呢？应该跳过0
func productExceptSelf(nums []int) []int {
	tp := 1
	zeroNum := 0
	for _, num := range nums {
		if num != 0 {
			tp *= num
			continue
		}
		zeroNum++
	}
	r := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			if zeroNum == 1 {
				r[i] = tp
			}
			continue
		}
		if zeroNum > 0 {
			continue
		}
		r[i] = tp / num
	}
	return r
}
