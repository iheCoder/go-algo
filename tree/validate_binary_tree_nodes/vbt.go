package validate_binary_tree_nodes

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	indeg := make([]int, n)

	for _, c := range leftChild {
		if c < 0 {
			continue
		}
		indeg[c]++
	}
	for _, c := range rightChild {
		if c < 0 {
			continue
		}
		indeg[c]++
	}

	root := -1
	for i, c := range indeg {
		if c != 0 {
			continue
		}

		if root != -1 {
			return false
		}
		root = i
	}
	if root == -1 {
		return false
	}

	queue := []int{root}
	seen := make(map[int]bool)
	seen[root] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, x := range []int{leftChild[p], rightChild[p]} {
			if x == -1 {
				continue
			}

			if seen[x] {
				return false
			}
			seen[x] = true
			queue = append(queue, x)
		}
	}

	return len(seen) == n
}
