package verify_preorder_serialization_of_a_binary_tree

import "strings"

func isValidSerialization(preorder string) bool {
	ps := strings.Split(preorder, ",")
	n := len(ps)
	var i int
	var f func() bool
	f = func() bool {
		if i >= n {
			return false
		}

		if ps[i] == "#" {
			i++
			return true
		}

		i++
		if !f() {
			return false
		}
		if !f() {
			return false
		}
		return true
	}

	return f() && i >= n
}
