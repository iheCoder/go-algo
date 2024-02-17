package longest_repeating_char_replace

func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	var l, r int
	var mc, ans int
	for r < len(s) {
		m[s[r]]++
		// 不理解为什么mc在左边界移动之后不需要更新呢？
		// 因为如果没有出现更大的，都不需要进行考虑了
		mc = maxInt(mc, m[s[r]])
		r++

		if r-l > mc+k {
			m[s[l]]--
			l++
		}
		ans = maxInt(ans, r-l)
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
