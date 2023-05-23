package palindrome_linked_list

import "testing"

type testData struct {
	nums     []int
	pos      int
	expected bool
}

func TestPDRevert(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{1},
			pos:      -1,
			expected: true,
		},
		{
			nums:     []int{1, 2},
			pos:      -1,
			expected: false,
		},
		{
			nums:     []int{1, 1},
			pos:      -1,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3},
			pos:      -1,
			expected: false,
		},
		{
			nums:     []int{1, 2, 1},
			pos:      -1,
			expected: true,
		},
		{
			nums:     []int{1, 2, 1, 2},
			pos:      -1,
			expected: false,
		},
	}

	for i, td := range tds {
		h := buildLinkedTableByArr(td.nums, td.pos)
		r := isPalindrome(h)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
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
