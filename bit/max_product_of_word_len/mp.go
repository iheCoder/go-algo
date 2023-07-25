package max_product_of_word_len

import "sort"

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) >= len(words[j])
	})

	reap := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		x := 0
		for j := 0; j < len(words[i]); j++ {
			x |= 1 << (words[i][j] - 'a')
		}
		reap[i] = x
	}

	var ans int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if reap[i]&reap[j] > 0 {
				continue
			}
			ans = maxInt(ans, len(words[i])*len(words[j]))
			break
		}
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
