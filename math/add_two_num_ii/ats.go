package add_two_num_ii

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1, n2 string
	p := l1
	for p != nil {
		n1 += strconv.Itoa(p.Val)
		p = p.Next
	}
	p = l2
	for p != nil {
		n2 += strconv.Itoa(p.Val)
		p = p.Next
	}
	if len(n1) < 19 && len(n2) < 19 {
		n1i, _ := strconv.Atoi(n1)
		n2i, _ := strconv.Atoi(n2)
		n := n1i + n2i
		cur := &ListNode{}
		var next *ListNode
		for n > 0 {
			cur = &ListNode{
				Val:  n % 10,
				Next: next,
			}
			n /= 10
			next = cur
		}
		return cur
	}
	ni := addNums(n1, n2)
	cur := &ListNode{}
	var next *ListNode
	for i := len(ni) - 1; i >= 0; i-- {
		cur = &ListNode{Val: int(ni[i]), Next: next}
		next = cur
	}
	return cur
}
func addNums(n1, n2 string) string {
	var r []byte
	i, j := len(n1)-1, len(n2)-1
	var carry byte
	var xn, yn byte
	for {
		if i < 0 && j < 0 && carry <= 0 {
			break
		}
		xn, yn = 0, 0
		if i >= 0 {
			xn = n1[i] - 48
		}
		if j >= 0 {
			yn = n2[j] - 48
		}
		x := xn + yn + carry
		r = append(r, x%10)
		carry = x / 10
		i--
		j--
	}
	reverseBytes(r)
	return string(r)
}

func reverseBytes(s []byte) []byte {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
