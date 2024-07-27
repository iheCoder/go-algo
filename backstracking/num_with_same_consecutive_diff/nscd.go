package num_with_same_consecutive_diff

func numsSameConsecDiff(n int, k int) []int {
	var ans []int
	var bt func(rem, cur, last int)
	bt = func(rem, num, last int) {
		if rem == 0 {
			ans = append(ans, num)
			return
		}

		less := last - k
		if less >= 0 {
			bt(rem-1, num*10+less, less)
		}
		more := last + k
		if more < 10 && less != more {
			bt(rem-1, num*10+more, more)
		}
	}

	for i := 1; i < 10; i++ {
		bt(n-1, i, i)
	}
	return ans
}
