package circular_arr_loop

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	m := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if m[i] == true {
			continue
		}

		slow, fast := i, getNextIndex(i, nums[i], n)
		m[i] = true
		for slow != fast {
			slow = getNextIndex(slow, nums[slow], n)
			m[slow] = true

			fast = getNextIndex(fast, nums[fast], n)
			fast = getNextIndex(fast, nums[fast], n)
		}

		fast = getNextIndex(fast, nums[fast], n)
		if slow == fast {
			continue
		}

		sv := nums[slow]
		for slow != fast {
			if sv*nums[fast] < 0 {
				break
			}
			fast = getNextIndex(fast, nums[fast], n)
		}
		if slow == fast {
			return true
		}
	}

	return false
}

func getNextIndex(i, offset, n int) int {
	r := (i + offset) % n
	if r >= 0 {
		return r
	}

	return r + n
}
