package jump_game

import "testing"

type jpTest struct {
	nums     []int
	expected bool
}

func TestJGFromRightToLeft(t *testing.T) {
	ts := []jpTest{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, jpt := range ts {
		r := canJumpRL(jpt.nums)
		if r != jpt.expected {
			t.Fatalf("nums %v expecte %v got %v", jpt.nums, jpt.expected, r)
		}
	}
}

func TestJGLR(t *testing.T) {
	ts := []jpTest{
		//{
		//	nums:     []int{2, 3, 1, 1, 4},
		//	expected: true,
		//},
		//{
		//	nums:     []int{3, 2, 1, 0, 4},
		//	expected: false,
		//},
		{
			nums:     []int{2, 5, 0, 0},
			expected: true,
		},
	}

	for _, jpt := range ts {
		r := canJumpLR(jpt.nums)
		if r != jpt.expected {
			t.Fatalf("nums %v expecte %v got %v", jpt.nums, jpt.expected, r)
		}
	}
}

func TestRightMost(t *testing.T) {
	ts := []jpTest{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			nums:     []int{2, 5, 0, 0},
			expected: true,
		},
	}

	for _, jpt := range ts {
		r := canJump(jpt.nums)
		if r != jpt.expected {
			t.Fatalf("nums %v expecte %v got %v", jpt.nums, jpt.expected, r)
		}
	}
}
