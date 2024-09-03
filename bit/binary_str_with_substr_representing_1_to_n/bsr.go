package binary_str_with_substr_representing_1_to_n

import "strings"

func help(s string, k, mi, ma int) bool {
	st := make(map[int]struct{})
	var t int
	for r := 0; r < len(s); r++ {
		// 将当前字符（'0' 或 '1'）转换为整数并添加到 t 的末尾
		t = (t << 1) + int(s[r]-'0')

		// 如果 r >= k，说明当前子串长度超过了 k，需要从 t 中减去最左边的位：
		if r >= k {
			t -= int(s[r-k]-'0') << k
		}

		if r >= k-1 && t >= mi && t <= ma {
			st[t] = struct{}{}
		}
	}

	return len(st) == ma-mi+1
}

func queryString(s string, n int) bool {
	if n == 1 {
		return strings.Contains(s, "1")
	}

	k := 30
	for (1 << k) >= n {
		k--
	}

	// (1<<(k-1)): 这是长度为 k-1 的二进制数的最大值。加 k-1 是为了确保字符串 s 足够长，以便容纳长度为 k 的所有子串
	// n - (1<<k) + k + 1: 这个表达式计算 s 是否足够长，以便容纳从 2^k 到 n 的所有二进制子串。n - (1<<k) 计算的是从 2^k 到 n 之间的距离，+ k + 1 确保字符串 s 有足够的字符来容纳这些长度为 k+1 的子串
	if len(s) < (1<<(k-1))+k-1 || len(s) < n-(1<<k)+k+1 {
		return false
	}

	return help(s, k, 1<<(k-1), (1<<k)-1) && help(s, k+1, 1<<k, n)
}
