package shortest_unsorted_continous_subarray

func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var i int
	for i < len(nums)-1 && nums[i] <= nums[i+1] {
		i++
	}
	if i == len(nums)-1 {
		return 0
	}

	mr := nums[i]
	for k := i + 1; k < len(nums); k++ {
		if nums[k] < mr {
			mr = nums[k]
		}
	}
	for i >= 0 {
		if nums[i] <= mr {
			break
		}
		i--
	}

	j := len(nums) - 1
	for j > 0 && nums[j] >= nums[j-1] {
		j--
	}
	mr = nums[j]
	for k := j - 1; k >= 0; k-- {
		if nums[k] > mr {
			mr = nums[k]
		}
	}
	for j < len(nums) {
		if nums[j] >= mr {
			break
		}
		j++
	}

	return j - i - 1
}
