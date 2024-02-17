package rm_dup_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var i, j int
	var temp int
	var count int
	temp = nums[j]
	count = 1
	j++
	i++
	for j < len(nums) {
		if temp == nums[j] {
			count++
		} else {
			count = 1
			temp = nums[j]
		}

		if count <= 2 {
			nums[i] = nums[j]
			i++
			j++
		} else {
			for temp == nums[j] {
				j++
				if j >= len(nums) {
					return i
				}
			}
		}
	}

	return i
}
