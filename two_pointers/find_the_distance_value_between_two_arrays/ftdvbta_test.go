package find_the_distance_value_between_two_arrays

import "testing"

type testData struct {
	arr1     []int
	arr2     []int
	d        int
	expected int
}

func TestSortLR(t *testing.T) {
	tds := []testData{
		{
			arr1:     []int{2, 1, 100, 3},
			arr2:     []int{-5, -2, 10, -3, 7},
			d:        6,
			expected: 1,
		},
		{
			arr1:     []int{-803, 715, -224, 909, 121, -296, 872, 807, 715, 407, 94, -8, 572, 90, -520, -867, 485, -918, -827, -728, -653, -659, 865, 102, -564, -452, 554, -320, 229, 36, 722, -478, -247, -307, -304, -767, -404, -519, 776, 933, 236, 596, 954, 464},
			arr2:     []int{817, 1, -723, 187, 128, 577, -787, -344, -920, -168, -851, -222, 773, 614, -699, 696, -744, -302, -766, 259, 203, 601, 896, -226, -844, 168, 126, -542, 159, -833, 950, -454, -253, 824, -395, 155, 94, 894, -766, -63, 836, -433, -780, 611, -907, 695, -395, -975, 256, 373, -971, -813, -154, -765, 691, 812, 617, -919, -616, -510, 608, 201, -138, -669, -764, -77, -658, 394, -506, -675, 523, 730, -790, -109, 865, 975, -226, 651, 987, 111, 862, 675, -398, 126, -482, 457, -24, -356, -795, -575, 335, -350, -919, -945, -979, 611},
			d:        37,
			expected: 0,
		},
	}

	for i, td := range tds {
		r := findTheDistanceValue(td.arr1, td.arr2, td.d)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

func TestAbs(t *testing.T) {
	t.Log(absInt(-1))
	t.Log(absInt(-10086))
	t.Log(absInt(1))
}

func TestBinarySearch(t *testing.T) {
	t.Log(binarySearch(1, []int{1}))
	t.Log(binarySearch(2, []int{1}))
	t.Log(binarySearch(2, []int{1, 2}))
	t.Log(binarySearch(2, []int{1, 3}))
	t.Log(binarySearch(2, []int{1, 4}))
	t.Log(binarySearch(2, []int{1, 3, 4}))
	t.Log(binarySearch(2, []int{1, 4, 5}))
}
