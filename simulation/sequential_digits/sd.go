package sequential_digits

import "sort"

func sequentialDigits(low int, high int) []int {
	var ans []int
	for i := 1; i <= 9; i++ {
		num := i
		for j := i + 1; j <= 9; j++ {
			num = num*10 + j
			if num >= low && num <= high {
				ans = append(ans, num)
			}
		}
	}

	sort.Ints(ans)
	return ans
}

func getNumWidth(num int) int {
	var width int
	for num > 0 {
		width++
		num /= 10
	}
	return width
}

func getDelta(width int) int {
	delta := 1
	for i := 0; i < width; i++ {
		delta = delta*10 + 1
	}

	return delta
}

func getMinSeqDigits(width int) int {
	var val int
	for i := 1; i <= width; i++ {
		val = val*10 + i
	}

	return width
}

func getMaxSeqDigits(width int) int {
	var val int
	for i := 1; i <= width; i++ {
		val = val*10 + (10 - i)
	}

	return val
}
