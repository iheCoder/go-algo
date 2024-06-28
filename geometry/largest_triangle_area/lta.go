package largest_triangle_area

import "math"

func largestTriangleArea(points [][]int) float64 {
	var ans float64
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				ans = math.Max(ans, triangleArea(points[i][0], points[j][0], points[k][0], points[i][1], points[j][1], points[k][1]))
			}
		}
	}

	return ans
}

func triangleArea(x1, x2, x3, y1, y2, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
}
