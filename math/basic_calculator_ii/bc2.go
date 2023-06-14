package basic_calculator_ii

func calculate(s string) int {
	stk := make([]int, 0)
	preSign := '+'
	var num int
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stk = append(stk, num)
			case '-':
				stk = append(stk, -num)
			case '*':
				stk[len(stk)-1] *= num
			case '/':
				stk[len(stk)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}

	var ans int
	for _, v := range stk {
		ans += v
	}
	return ans
}
