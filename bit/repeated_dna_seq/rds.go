package repeated_dna_seq

func findRepeatedDnaSequencesHashForce(s string) []string {
	if len(s) < 10 {
		return nil
	}

	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		m[s[i:i+10]]++
	}

	ans := make([]string, 0)
	for sb, c := range m {
		if c > 1 {
			ans = append(ans, sb)
		}
	}

	return ans
}

const L = 10

var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequencesHashSlidingWindow(s string) []string {
	if len(s) < L {
		return nil
	}

	n := len(s)
	var ans []string
	// 将前L长度编码到数字x中
	var x int
	for _, ch := range s[:L-1] {
		x = x<<2 | bin[byte(ch)]
	}
	cnt := make(map[int]int)
	for i := 0; i <= n-L; i++ {
		// 每进来一个DNA，移出最前面的，并添加最新的
		x = (x<<2 | bin[s[i+L-1]]) & (1<<(L<<1) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+L])
		}
	}
	return ans
}
