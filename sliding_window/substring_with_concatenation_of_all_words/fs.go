package substring_with_concatenation_of_all_words

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {
	if len(words)*len(words[0]) > len(s) {
		return nil
	}

}

func computeLPS(pattern string) []int {
	m := len(pattern)
	next := make([]int, 1, m)
	r := 1
	l := 0
	for r < m {
		if pattern[l] == pattern[r] {
			l++
			r++
			next = append(next, l)
		} else if l != 0 {
			l = next[l-1]
		} else {
			next = append(next, 0)
			r++
		}
	}

	return next
}
