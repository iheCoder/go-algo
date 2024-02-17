package add_two_numbers

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dh := &ListNode{}
	prev := dh
	var l *ListNode
	var carry, x int
	var l1v, l2v int
	for l1 != nil || l2 != nil || carry > 0 {
		l1v, l2v = 0, 0
		if l1 != nil {
			l1v = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2v = l2.Val
			l2 = l2.Next
		}
		x = l1v + l2v + carry
		carry = x / 10
		x %= 10
		l = &ListNode{Val: x}
		prev.Next = l
		prev = l
	}
	return dh.Next
}

// 最简单的是取出来，加之后，再放进去
// 超出限制了！！！这可以是100位啊
func addTwoNumbersFetchAddThenConstruct(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1, n2 int
	p := l1
	var i int
	for p != nil {
		n1 += p.Val * pow10(i)
		p = p.Next
		i++
	}
	p = l2
	i = 0
	for p != nil {
		n2 += p.Val * pow10(i)
		p = p.Next
		i++
	}
	n := n1 + n2
	root := &ListNode{}
	if n == 0 {
		return root
	}
	root.Val = n % 10
	n /= 10
	prev := root
	var cur *ListNode
	for n > 0 {
		cur = &ListNode{
			Val: n % 10,
		}
		prev.Next = cur
		prev = cur
		n /= 10
	}
	return root
}

func pow10(x int) int {
	return int(math.Pow10(x))
}
