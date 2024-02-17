package construct_quad_tree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	if len(grid) == 1 {
		return &Node{
			Val:    grid[0][0] == 1,
			IsLeaf: true,
		}
	}

	alterGrid(grid, 0, 0, len(grid))

	return cqt(grid, 0, 0, len(grid))
}

func cqt(grid [][]int, i, j, n int) *Node {
	if n == 1 || grid[i][j] == n {
		val := grid[i][j] == 1
		if n > 1 {
			val = grid[i][j+1] == 1
		}
		return &Node{Val: val, IsLeaf: true}
	}
	root := &Node{Val: true, IsLeaf: false}

	hn := n >> 1
	root.TopLeft = cqt(grid, i, j, hn)
	root.TopRight = cqt(grid, i, j+hn, hn)
	root.BottomLeft = cqt(grid, i+hn, j, hn)
	root.BottomRight = cqt(grid, i+hn, j+hn, hn)
	return root
}

func alterGrid(grid [][]int, i, j, n int) int {
	if n == 1 {
		return grid[i][j]
	}

	hn := n >> 1
	tl := alterGrid(grid, i, j, hn)
	tr := alterGrid(grid, i, j+hn, hn)
	bl := alterGrid(grid, i+hn, j, hn)
	br := alterGrid(grid, i+hn, j+hn, hn)
	if tl == -1 || tr == -1 || bl == -1 || br == -1 {
		return -1
	}
	if tl == tr && bl == br && tl == bl {
		grid[i][j] = n
		return grid[i][j+1]
	}
	return -1
}
