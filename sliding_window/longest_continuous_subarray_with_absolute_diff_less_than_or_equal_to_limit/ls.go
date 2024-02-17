package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

func longestSubarray(nums []int, limit int) int {
	if len(nums) == 1 {
		return 1
	}

	// 1. 初始化，找到前两个元素的最大、最小。若差大于limit，那么就一直移动
	i, j := 0, 1
	for j < len(nums) && !diffSatisfy(nums[i], nums[j], limit) {
		i++
		j++
	}
	if j >= len(nums) {
		return 1
	}
	r := 2
	// 2. push新元素。若新元素大于最大值，则判断是否差超过limit，超过就移出最大值以及最大值之前的元素
	arr := make([]int, 0)
	push := func(k int) {
		arr = append(arr, k)
		l := len(arr) - 1
		for l > 0 && nums[arr[l]] > nums[arr[l-1]] {
			arr[l], arr[l-1] = arr[l-1], arr[l]
			l--
		}
	}
	push(i)
	push(j)
	j++
	for j < len(nums) {
		if nums[j] > nums[arr[0]] {
			for len(arr) > 0 && !diffSatisfy(nums[j], nums[arr[len(arr)-1]], limit) {
				// 前面的index也有可能小于后面的，所以取最大值，反正这区间的都不能要
				i = maxInt(i, arr[len(arr)-1]+1)
				arr = arr[:len(arr)-1]
			}
		} else if nums[j] < nums[arr[len(arr)-1]] {
			for len(arr) > 0 && !diffSatisfy(nums[arr[0]], nums[j], limit) {
				i = maxInt(i, arr[0]+1)
				arr = arr[1:]
			}
		}
		if len(arr) > 0 {
			r = maxInt(r, j-i+1)
		} else if len(arr) == 0 {
			i = j
		}
		push(j)
		j++
	}

	// 3. 如此重复第二步，直到得到最大值
	return r
}

func diffSatisfy(x, y, limit int) bool {
	diff := x - y
	if diff < 0 {
		diff = -diff
	}
	return diff <= limit
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
