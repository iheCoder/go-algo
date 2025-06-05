package max_num_of_occurences_of_a_substring

const mod int = 1000000007

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	n := len(s)
	occ := make(map[int]int)
	exist := make(map[byte]int)
	ans, existCnt := 0, 0
	rabin, base, baseMul := 0, 26, 26

	for i := 0; i < minSize-1; i++ {
		exist[s[i]]++
		if exist[s[i]] == 1 {
			existCnt++
		}

		rabin = (rabin*base + int(s[i]-'a')) % mod
		baseMul = baseMul * base % mod
	}

	for i := minSize - 1; i < n; i++ {
		exist[s[i]]++
		if exist[s[i]] == 1 {
			existCnt++
		}
		rabin = (rabin*base + int(s[i]-'a')) % mod

		if i >= minSize {
			exist[s[i-minSize]]--
			if exist[s[i-minSize]] == 0 {
				existCnt--
			}
			rabin = (rabin - baseMul*int(s[i-minSize]-'a')%mod + mod) % mod
		}

		if existCnt <= maxLetters {
			occ[rabin]++
			if occ[rabin] > ans {
				ans = occ[rabin]
			}
		}

	}

	return ans
}

func maxFreqAn(s string, maxLetters int, minSize int, maxSize int) int {
	n := len(s)
	occ := make(map[string]int)
	var ans int

	for i := 0; i < n-minSize+1; i++ {
		x := s[i : i+minSize]

		em := make(map[byte]struct{})
		for j := range x {
			em[x[j]] = struct{}{}
		}

		if len(em) <= maxLetters {
			occ[x]++
			if occ[x] > ans {
				ans = occ[x]
			}
		}
	}

	return ans
}
