package gray_code

func grayCode(n int) []int {
	ans := make([]int, 1, 1<<n)
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|1<<(i-1))
		}
	}
	return ans
}

// 设i为正整数，^为亦或符号，函数f(i)=i^(i>>1)。求证f(i)与f(i+1)的二进制表示 恰好一位不同
