package reorder_list

import "testing"

type testData struct {
	nums     []int
	pos      int
	expected []int
}

func TestRL(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{1, 2, 3, 4},
			pos:      -1,
			expected: []int{1, 4, 2, 3},
		},
		{
			nums:     []int{1, 2, 3},
			pos:      -1,
			expected: []int{1, 3, 2},
		},
		{
			nums:     []int{1, 2},
			pos:      -1,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1},
			pos:      -1,
			expected: []int{1},
		},
	}

	for i, td := range tds {
		h := buildLinkedTableByArr(td.nums, td.pos)
		reorderList(h)
		if !checkListArrEqual(h, td.expected) {
			t.Fatalf("index %d expect %v got %v", i, td.expected, h)
		}
	}
}

func checkListArrEqual(h *ListNode, arr []int) bool {
	if h == nil && len(arr) == 0 {
		return true
	}

	p := h
	for i := 0; i < len(arr); i++ {
		if p == nil {
			return false
		}
		if p.Val != arr[i] {
			return false
		}
		p = p.Next
	}

	return p == nil
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
