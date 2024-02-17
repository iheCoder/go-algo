package total_hamming_distance

func totalHammingDistanceXor(nums []int) int {
	n := len(nums)
	var total int
	for i := 0; i < n-1; i++ {
		a := nums[i]
		for j := i + 1; j < n; j++ {
			total += bitOneCount(a ^ nums[j])
		}
	}
	return total
}

func bitOneCount(n int) int {
	var count int
	for n > 0 {
		count++
		n &= (n - 1)
	}
	return count
}

func totalHammingDistance(nums []int) int {
	n := len(nums)
	var total int
	for i := 0; i < 30; i++ {
		var c int
		for _, num := range nums {
			c += num >> i & 1
		}
		total += c * (n - c)
	}
	return total
}
