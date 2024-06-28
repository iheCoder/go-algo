package score_of_parentheses

func scoreOfParentheses(s string) int {
	n := len(s)
	if n == 2 {
		return 1
	}

	var bar, i int
	for i < n {
		if s[i] == '(' {
			bar++
		} else {
			bar--
			if bar == 0 {
				if i == n-1 {
					return 2 * scoreOfParentheses(s[1:n-1])
				}
				return scoreOfParentheses(s[:i+1]) + scoreOfParentheses(s[i+1:])
			}
		}
		i++
	}
	return 0
}
