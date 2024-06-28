package validate_stack_seq

import "testing"

func TestVSS(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	poped := []int{4, 5, 3, 2, 1}
	t.Log(validateStackSequences(pushed, poped))

	pushed = []int{1, 0}
	poped = []int{1, 0}
	t.Log(validateStackSequences(pushed, poped))
}
