package num_of_substrings_containing_all_three_char

func numberOfSubstrings(s string) int {
	n, ans := len(s), 0
	l, r := 0, -1
	var cnt [3]int

	for l < n {
		for r+1 < n && !(cnt[0] >= 1 && cnt[1] >= 1 && cnt[2] >= 1) {
			r++
			cnt[s[r]-'a']++
		}

		if cnt[0] >= 1 && cnt[1] >= 1 && cnt[2] >= 1 {
			ans += n - r
		}

		cnt[s[l]-'a']--
		l++
	}

	return ans
}
