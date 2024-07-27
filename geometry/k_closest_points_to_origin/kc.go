package k_closest_points_to_origin

import "sort"

func kClosest(points [][]int, k int) [][]int {
	type pair struct {
		i    int
		path int
	}
	ps := make([]*pair, 0, len(points))
	for i, point := range points {
		ps = append(ps, &pair{
			i:    i,
			path: calPath(point[0], point[1]),
		})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].path <= ps[j].path
	})

	ans := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, points[ps[i].i])
	}
	return ans
}

func calPath(x, y int) int {
	return x*x + y*y
}

// func less(p, q []int) bool {
//    return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
//}
//
//func kClosest(points [][]int, k int) (ans [][]int) {
//    rand.Shuffle(len(points), func(i, j int) {
//        points[i], points[j] = points[j], points[i]
//    })
//
//    var quickSelect func(left, right int)
//    quickSelect = func(left, right int) {
//        if left == right {
//            return
//        }
//        pivot := points[right] // 取当前区间 [left,right] 最右侧元素作为切分参照
//        lessCount := left
//        for i := left; i < right; i++ {
//            if less(points[i], pivot) {
//                points[i], points[lessCount] = points[lessCount], points[i]
//                lessCount++
//            }
//        }
//        // 循环结束后，有 lessCount 个元素比 pivot 小
//        // 把 pivot 交换到 points[lessCount] 的位置
//        // 交换之后，points[lessCount] 左侧的元素均小于 pivot，points[lessCount] 右侧的元素均不小于 pivot
//        points[right], points[lessCount] = points[lessCount], points[right]
//        if lessCount+1 == k {
//            return
//        } else if lessCount+1 < k {
//            quickSelect(lessCount+1, right)
//        } else {
//            quickSelect(left, lessCount-1)
//        }
//    }
//    quickSelect(0, len(points)-1)
//    return points[:k]
//}
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/k-closest-points-to-origin/solutions/477916/zui-jie-jin-yuan-dian-de-k-ge-dian-by-leetcode-sol/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
