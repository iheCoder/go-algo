package find_2_non_overlapping_sub_arr_each_with_target_sum

import "testing"

func TestMinSumOfLengths(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{
			name:   "two single element windows",
			arr:    []int{3, 2, 2, 4, 3},
			target: 3,
			want:   2,
		},
		{
			name:   "choose shortest non-overlapping pair",
			arr:    []int{7, 3, 4, 7},
			target: 7,
			want:   2,
		},
		{
			name:   "only one target window",
			arr:    []int{4, 3, 2, 6, 2, 3, 4},
			target: 6,
			want:   -1,
		},
		{
			name:   "adjacent windows are allowed",
			arr:    []int{1, 1, 1, 1, 1},
			target: 2,
			want:   4,
		},
		{
			name:   "left side keeps shortest historical window",
			arr:    []int{1, 2, 1, 2, 1, 2, 1},
			target: 3,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSumOfLengths(tt.arr, tt.target); got != tt.want {
				t.Fatalf("minSumOfLengths(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}
