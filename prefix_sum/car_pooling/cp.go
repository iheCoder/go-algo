package car_pooling

func carPooling(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return true
	}

	var toMax int
	for _, trip := range trips {
		toMax = maxInt(toMax, trip[2])
	}

	pass := make([]int, toMax+1)
	for _, trip := range trips {
		pass[trip[1]] += trip[0]
		pass[trip[2]] -= trip[0]
	}

	var count int
	for _, p := range pass {
		count += p
		if count > capacity {
			return false
		}
	}

	return true
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// func carPooling(trips [][]int, capacity int) bool {
//	if len(trips) == 0 {
//		return true
//	}
//
//	// 按左端点排序
//	sort.Slice(trips, func(i, j int) bool {
//		return trips[i][1] < trips[j][1]
//	})
//
//	merged := make([][]int, 0, len(trips))
//	merged = append(merged, trips[0])
//	for i := 1; i < len(trips); i++ {
//		last := merged[len(merged)-1]
//		if last[2] > trips[i][1] {
//			last[2] = maxInt(last[2], trips[i][2])
//			last[0] += trips[i][0]
//		} else {
//			merged = append(merged, trips[i])
//		}
//	}
//
//	for _, trip := range merged {
//		if trip[0] > capacity {
//			return false
//		}
//	}
//
//	return true
//}
