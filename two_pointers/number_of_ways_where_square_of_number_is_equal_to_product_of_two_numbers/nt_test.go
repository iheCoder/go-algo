package number_of_ways_where_square_of_number_is_equal_to_product_of_two_numbers

import "testing"

type testData struct {
	nums1    []int
	nums2    []int
	expected int
}

func TestNT(t *testing.T) {
	tds := []testData{
		//{
		//	nums1:    []int{1, 1},
		//	nums2:    []int{1, 1, 1},
		//	expected: 9,
		//},
		{
			nums1:    []int{7, 7, 8, 3},
			nums2:    []int{1, 2, 9, 7},
			expected: 2,
		},
	}

	for i, td := range tds {
		r := numTriplets(td.nums1, td.nums2)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}
