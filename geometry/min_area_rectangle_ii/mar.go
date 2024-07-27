package min_area_rectangle_ii

import "math"

type point struct {
	x, y int
}

func minAreaFreeRect(points [][]int) float64 {
	n := len(points)
	a := make([]point, n)
	for i := 0; i < n; i++ {
		a[i] = point{
			x: points[i][0],
			y: points[i][1],
		}
	}

	seen := make(map[int]map[point][]point)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			center := point{
				x: a[i].x + a[j].x,
				y: a[i].y + a[j].y,
			}
			r2 := (a[i].x-a[j].x)*(a[i].x-a[j].x) + (a[i].y-a[j].y)*(a[i].y-a[j].y)
			if _, ok := seen[r2]; !ok {
				seen[r2] = make(map[point][]point)
			}
			seen[r2][center] = append(seen[r2][center], a[i])
		}
	}

	ans := math.MaxFloat64
	for _, info := range seen {
		for center, cs := range info {
			cl := len(cs)
			for i := 0; i < cl; i++ {
				for j := i + 1; j < cl; j++ {
					p, q := cs[i], cs[j]

					// 	•	Q 点是通过 center 对称 q 得到的点。
					//	•	公式为：Q.x = center.x - (q.x - center.x) 和 Q.y = center.y - (q.y - center.y)。
					//	•	简化为：Q.x = center.x - q.x 和 Q.y = center.y - q.y。
					Q := point{
						center.x - q.x,
						center.y - q.y,
					}

					// 求矩形两边长度并相乘得到面积
					area := distance(p, q) * distance(p, Q)
					if area < ans {
						ans = area
					}
				}
			}
		}
	}

	if ans < math.MaxFloat64 {
		return ans
	}

	return 0
}

func distance(p1, p2 point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}
