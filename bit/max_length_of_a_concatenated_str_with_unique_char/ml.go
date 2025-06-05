package max_length_of_a_concatenated_str_with_unique_char

func maxLength(arr []string) int {
	n := len(arr)
	barr := make([]uint32, 0, n)
	for i := 0; i < n; i++ {
		var ba uint32
		reap := false
		for j := 0; j < len(arr[i]); j++ {
			d := uint32(1 << (arr[i][j] - 'a'))
			if d&ba > 0 {
				reap = true
				break
			}
			ba |= d
		}
		if reap {
			continue
		}
		barr = append(barr, ba)
	}

	n = len(barr)
	var ans uint32
	var bs func(int, uint32)
	bs = func(pos int, mask uint32) {
		if pos >= n {
			ans = maxInt(ans, countOnes(mask))
			return
		}

		if barr[pos]&mask == 0 {
			bs(pos+1, barr[pos]|mask)
		}
		bs(pos+1, mask)
	}
	bs(0, 0)
	return int(ans)
}

func maxInt(x, y uint32) uint32 {
	if x > y {
		return x
	}

	return y
}

func countOnes(x uint32) uint32 {
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x & 0x0F0F0F0F) + ((x >> 4) & 0x0F0F0F0F)
	x = (x & 0x00FF00FF) + ((x >> 8) & 0x00FF00FF)
	x = (x & 0x0000FFFF) + ((x >> 16) & 0x0000FFFF)
	return x
}
