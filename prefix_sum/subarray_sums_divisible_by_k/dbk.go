package subarray_sums_divisible_by_k

func subarraysDivByK(nums []int, k int) int {
	var sum int
	m := map[int]int{0: 1}
	ans := 0
	for _, num := range nums {
		sum += num
		mod := (sum%k + k) % k
		ans += m[mod]
		m[mod]++
	}

	return ans
}

// 单纯的前缀和当然是不行的，超出了时间
func subarraysDivByKPrefixSum(nums []int, k int) int {
	var sum int
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sums[i+1] = sum
	}

	var ans int
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if (sums[i]-sums[j])%k == 0 {
				ans++
			}
		}
	}
	return ans
}
