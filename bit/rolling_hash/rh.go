package rolling_hash

const (
	B   = 31         // 进制
	MOD = 1000000007 // 取模
)

func rollingHash(s string, L int) []int {
	n := len(s)
	if n < L {
		return nil
	}

	hashes := make([]int, n-L+1)
	hash, power := 0, 1

	// 计算第一个子串的hash值
	for i := 0; i < L; i++ {
		hash = (hash*B + int(s[i]-'a'+1)) % MOD
		if i < L-1 {
			power = (power * B) % MOD
		}
	}

	// 滚动计算后续子串的hash值
	for i := 0; i <= n-L; i++ {
		hash = ((hash-int(s[i-1]-'a'+1)*power%MOD+MOD)%MOD*B + int(s[i+L-1]-'a'+1)) % MOD
		hashes[i] = hash
	}

	return hashes
}
