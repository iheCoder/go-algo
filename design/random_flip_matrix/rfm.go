package random_flip_matrix

import (
	"math/rand"
)

type Solution struct {
	m, n  int
	total int
	mp    map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{
		m:     m,
		n:     n,
		total: m * n,
		mp:    make(map[int]int),
	}
}

func (s *Solution) Flip() []int {
	// 选取随机数
	x := rand.Intn(s.total)
	s.total--

	var ans []int
	// 若选择数被修改了，那么就选择储存的值
	// 储存的值一定没有被选择，因为储存的值一定来自于没有被储存的最后数
	if y, ok := s.mp[x]; ok {
		ans = []int{y / s.n, y % s.n}
	} else {
		ans = []int{x / s.n, x % s.n}
	}

	// 若最后数被修改了，那么修改随机数值设置为之前最后数的值
	// 否则，随机数值设置为最后树
	if y, ok := s.mp[s.total]; ok {
		s.mp[x] = y
	} else {
		s.mp[x] = s.total
	}

	return ans
}

func (s *Solution) Reset() {
	s.total = s.m * s.n
	s.mp = make(map[int]int)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(matrix, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
// type Solution struct {
//	matrix [][]int
//	m, n   int
//	occupy int
//}
//
//func Constructor(m int, n int) Solution {
//	matrix := make([][]int, m)
//	for i := 0; i < m; i++ {
//		matrix[i] = make([]int, n)
//	}
//	return Solution{
//		matrix: matrix,
//		m:      m,
//		n:      n,
//	}
//}
//
//func (s *Solution) Flip() []int {
//	var c int
//	var x []int
//	rem := s.m*s.n - s.occupy
//	target := rand.Intn(rem)
//	defer func() {
//		s.occupy++
//		s.matrix[x[0]][x[1]] = 1
//	}()
//	for i := 0; i < s.m; i++ {
//		for j := 0; j < s.n; j++ {
//			if s.matrix[i][j] == 0 {
//				if target == c {
//					x = []int{i, j}
//					return x
//				}
//				c++
//			}
//		}
//	}
//
//	return x
//}
//
//func (s *Solution) Reset() {
//	for i := 0; i < s.m; i++ {
//		for j := 0; j < s.n; j++ {
//			s.matrix[i][j] = 0
//		}
//	}
//	s.occupy = 0
//}
