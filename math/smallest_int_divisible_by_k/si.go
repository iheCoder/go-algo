package smallest_int_divisible_by_k

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	cur, res := 0, 1
	for {
		cur = (cur*10 + 1) % k
		if cur == 0 {
			return res
		}
		res++
	}
}
