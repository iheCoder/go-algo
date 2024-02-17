package evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	stk := make([]int, 0, len(tokens))
	var x, y int
	for _, token := range tokens {
		switch token {
		case "+":
			y = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			x = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk = append(stk, x+y)
		case "-":
			y = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			x = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk = append(stk, x-y)
		case "*":
			y = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			x = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk = append(stk, x*y)
		case "/":
			y = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			x = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			stk = append(stk, x/y)
		default:
			tmp, _ := strconv.Atoi(token)
			stk = append(stk, tmp)
		}
	}
	return stk[0]
}
