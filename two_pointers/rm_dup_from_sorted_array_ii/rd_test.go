package rm_dup_from_sorted_array_ii

import "testing"

type testData struct {
	nums         []int
	expected     int
	expectedNums []int
}

func TestRD(t *testing.T) {
	tds := []testData{
		{
			nums:         []int{1, 1, 1, 2, 2, 3},
			expected:     5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
	}

	for i, td := range tds {
		removeDuplicates(td.nums)
		if !checkExpected(td.nums, td.expectedNums) {
			t.Fatalf("index %d expect %v got %v", i, td.expectedNums, td.nums)
		}
	}
}

func checkExpected(nums, eNums []int) bool {
	n := len(eNums)
	if len(nums) < n {
		return false
	}

	for i := 0; i < n; i++ {
		if nums[i] != eNums[i] {
			return false
		}
	}

	return true
}
