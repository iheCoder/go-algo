package multiply_strings

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 先计算每位的值
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}

	// 将每位超过10的重新进位计算
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}

	// 组合得到答案
	var ans string
	for i := 0; i < len(ansArr); i++ {
		ans += strconv.Itoa(ansArr[i])
	}
	for len(ans) > 0 && ans[0] == '0' {
		ans = ans[1:]
	}
	return ans
}
