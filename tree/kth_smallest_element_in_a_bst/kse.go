package kth_smallest_element_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var c int
	n := root
	var p *TreeNode
	stk := make([]*TreeNode, 0)
	for n != nil || len(stk) > 0 {
		if n != nil {
			stk = append(stk, n)
			n = n.Left
		} else {
			p = stk[len(stk)-1]
			c++
			if c == k {
				return p.Val
			}
			stk = stk[:len(stk)-1]
			n = p.Right
		}
	}
	return 0
}

type MyBst struct {
	root    *TreeNode
	nodeNum map[*TreeNode]int // 统计以每个结点为根结点的子树的结点数，并存储在哈希表中
}

// 统计以 node 为根结点的子树的结点数
func (t *MyBst) countNodeNum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	t.nodeNum[node] = 1 + t.countNodeNum(node.Left) + t.countNodeNum(node.Right)
	return t.nodeNum[node]
}

// 返回二叉搜索树中第 k 小的元素
func (t *MyBst) kthSmallest(k int) int {
	node := t.root
	for {
		leftNodeNum := t.nodeNum[node.Left]
		if leftNodeNum < k-1 {
			node = node.Right
			k -= leftNodeNum + 1
		} else if leftNodeNum == k-1 {
			return node.Val
		} else {
			node = node.Left
		}
	}
}

func kthSmallestPre(root *TreeNode, k int) int {
	t := &MyBst{root, map[*TreeNode]int{}}
	t.countNodeNum(root)
	return t.kthSmallest(k)
}
