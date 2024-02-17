package array_nesting

func arrayNesting(nums []int) int {
	n := len(nums)
	visited := make([]int, n)
	var ans int
	for i := 0; i < n; i++ {
		if visited[i] > 0 {
			continue
		}
		j := nums[i]
		visited[i] = 1
		cur := 1
		for visited[j] == 0 {
			visited[j] = 1
			j = nums[j]
			cur++
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
