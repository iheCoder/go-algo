package count_num_of_nice_subarrays

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	odd := make([]int, n+1)
	var cnt int
	for i, num := range nums {
		if num&1 > 0 {
			cnt++
			odd[cnt] = i
		}
	}
	cnt++
	odd[0], odd[cnt] = -1, n

	var ans int
	for i := 1; i <= cnt-k; i++ {
		ans += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
	}
	return ans
}
