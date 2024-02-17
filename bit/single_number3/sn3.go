package single_number3

func singleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	lb := xor & (-xor)
	var one, two int
	for _, num := range nums {
		if num&lb == 0 {
			one ^= num
		} else {
			two ^= num
		}
	}

	return []int{one, two}
}
