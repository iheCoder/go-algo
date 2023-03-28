package find_the_dup_number

import "testing"

type testData struct {
	nums     []int
	expected int
}

func TestFDFactorial(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 4, 2, 3, 5, 3, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 1, 1},
			expected: 1,
		},
		{
			// 这个测试数据就宣告了阶乘方法失效
			nums:     []int{3, 2, 2, 2, 4},
			expected: 2,
		},
	}

	for i, td := range tds {
		r := findDuplicateFac(td.nums)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestTPAsIndex(t *testing.T) {
	tds := []testData{
		{
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 4, 2, 3, 5, 3, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 1, 1},
			expected: 1,
		},
		{
			nums:     []int{3, 2, 2, 2, 4},
			expected: 2,
		},
		{
			nums:     []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			expected: 9,
		},
	}

	for i, td := range tds {
		r := findDuplicate(td.nums)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
