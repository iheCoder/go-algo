package validate_stack_seq

func validateStackSequences(pushed []int, popped []int) bool {
	stk := []int{pushed[0]}
	i, j := 0, 0

	for {
		stk = append(stk, pushed[i])
		i++

		for len(stk) > 0 && stk[len(stk)-1] == popped[j] {
			j++
			stk = stk[:len(stk)-1]
			if j >= len(popped) {
				return true
			}
		}

		if i >= len(pushed) {
			return false
		}
	}

	return false
}
