package linked_list_components

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	p := head
	var ans int
	for p != nil {
		if m[p.Val] {
			ans++
			for p != nil && m[p.Val] {
				p = p.Next
			}
		}
		if p != nil {
			p = p.Next
		}
	}
	return ans
}
