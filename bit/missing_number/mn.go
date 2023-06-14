package missing_number

func missingNumberMath(nums []int) int {
	n := len(nums)
	total := n * (n + 1) >> 1
	var sum int
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	return total - sum
}

func missingNumber(nums []int) int {
	var ans int
	for i, num := range nums {
		ans ^= num ^ i
	}
	return ans ^ len(nums)
}
