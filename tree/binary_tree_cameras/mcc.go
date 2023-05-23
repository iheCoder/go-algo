package binary_tree_cameras

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 我的思路其实是0代表子树为空，1代表子树开监控了，-1代表子树需要开监控了。那么在子树都为空的情况下就返回需要开监控，在子树存在需要开监控的情况下就开监控，并返回1；其他情况下就返回0，代表当前节点被监控覆盖
// 并且还要注意特殊情况就是root也为-1需要开监控
func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans int
	var mcc func(n *TreeNode) int
	mcc = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		lon := mcc(n.Left)
		ron := mcc(n.Right)
		if lon == 0 && ron == 0 {
			return -1
		}
		if lon < 0 || ron < 0 {
			ans++
			return 1
		}
		return 0
	}
	if mcc(root) < 0 {
		ans++
	}
	return ans
}
