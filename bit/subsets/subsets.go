package subsets

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	n := len(nums)
	var ans [][]int
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i := 0; i < n; i++ {
			if mask>>i&1 > 0 {
				set = append(set, nums[i])
			}
		}
		ans = append(ans, set)
	}

	return ans
}

func subsetsBackTrack(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
