package linked_list_cycle_ii

import "testing"

type testData struct {
	nums     []int
	pos      int
	expected int
}

func TestTPABCL(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{3, 2, 0, -4},
			pos:      1,
			expected: 1,
		},
		{
			nums:     []int{1, 2},
			pos:      0,
			expected: 0,
		},
		{
			nums:     []int{1},
			pos:      -1,
			expected: -1,
		},
	}

	for i, td := range tds {
		h := buildLinkedTableByArr(td.nums, td.pos)
		r := detectCycle(h)
		if r != getItemByIndex(h, td.expected) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func getItemByIndex(head *ListNode, pos int) *ListNode {
	if pos < 0 {
		return nil
	}

	p := head
	for p != nil && pos > 0 {
		p = p.Next
		pos--
	}
	return p
}

func buildLinkedTableByArr(arr []int, index int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	var start *ListNode
	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	tail := head
	pre := head
	for i := 1; i < len(arr); i++ {
		tail = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		pre.Next = tail
		if i == index {
			start = tail
		}
		pre = tail
	}
	if start != nil {
		tail.Next = start
	} else if index == 0 {
		tail.Next = head
	}
	return head
}
