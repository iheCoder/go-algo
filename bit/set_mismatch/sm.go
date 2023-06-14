package set_mismatch

func findErrorNums(nums []int) []int {
	var xor int
	for i, num := range nums {
		xor ^= (i + 1) ^ num
	}

	lb := xor & (-xor)
	var n1, n2 int
	for i, num := range nums {
		if lb&num == 0 {
			n1 ^= num
		} else {
			n2 ^= num
		}

		if lb&(i+1) == 0 {
			n1 ^= i + 1
		} else {
			n2 ^= i + 1
		}
	}

	for _, num := range nums {
		if num == n1 {
			return []int{n1, n2}
		}
	}
	return []int{n2, n1}
}

// 本想通过x^y = m, x-y = n两个等式求出来，但是显然不可能
// 	var sum int
//	l := len(nums)
//	total := l * (l + 1) / 2
//	var m int
//	for i, num := range nums {
//		m ^= (i + 1) ^ num
//		sum += num
//	}
//	n := total - sum
//	y := n ^ m
//	x := n + y
//	return []int{y, x}
