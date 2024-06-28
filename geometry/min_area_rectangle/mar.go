package min_area_rectangle

import "math"

func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}

	ps := make(map[int]struct{})
	for _, point := range points {
		ps[getKey(point[0], point[1])] = struct{}{}
	}

	ans := math.MaxInt
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] != points[j][0] && points[i][1] != points[j][1] {
				if contains(ps, getKey(points[i][0], points[j][1])) && contains(ps, getKey(points[j][0], points[i][1])) {
					ans = minInt(ans, abs(points[i][0]-points[j][0])*abs(points[i][1]-points[j][1]))
				}
			}
		}
	}

	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func contains(s map[int]struct{}, i int) bool {
	_, ok := s[i]
	return ok
}

func getKey(x, y int) int {
	return 40001*x + y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
