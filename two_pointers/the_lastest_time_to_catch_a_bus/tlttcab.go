package the_lastest_time_to_catch_a_bus

import "sort"

func latestTimeCatchTheBusAllpossible(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	pm := make(map[int]bool)
	for _, passenger := range passengers {
		pm[passenger] = true
	}

	arrives := make([]int, 0)
	var i, j int
	var count int
	for i < len(buses) && j < len(passengers) {
		if passengers[j] <= buses[i] {
			count++
			at := passengers[j] - 1
			for at <= buses[i] {
				if !pm[at] {
					arrives = append(arrives, at)
					break
				}
				at--
			}
			j++
		} else {
			i++
			count = 0
		}
		if count >= capacity {
			i++
			count = 0
		}
	}

	// 前者是最后一辆车没有满。可能后面的乘客都晚于最后一辆，那么j肯定小于lp，count可能为0也可能小于cap
	// 或者后面没有乘客了，j等于lp，count也可能为0或者小于cap
	// 后者是根本没有到最后一辆，乘客就没了
	// 总之
	if (count < capacity && (i == len(buses))) || (i < len(buses)) {
		return buses[len(buses)-1]
	}

	return arrives[len(arrives)-1]
}

// 与其记录什么乘客在什么公交上不如记录所有可能上的位置
// 我认为应该提前终止这种方式，因为空间复杂度太高，代码写的太复杂，应该会有更优雅的方法
//taken := make([]map[int][]int, len(buses))

//for k := 0; k < capacity; k++ {
//	for j < len(passengers) && passengers[j] > buses[i] {
//		j++
//	}
//	if j >= len(passengers) {
//		arrives = append(arrives, buses[len(buses)-1])
//		break Loop
//	}
//	at := passengers[j] - 1
//	for at <= buses[i] {
//		if !pm[at] {
//			arrives = append(arrives, at)
//			break
//		}
//		at--
//	}
//	j++
//}
//i++

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	var i, j, c int
	for ; i < len(buses); i++ {
		// 关键还要有一个判断c是否为空的条件
		for c = capacity; j < len(passengers) && passengers[j] <= buses[i] && c > 0; c-- {
			j++
		}
	}

	var ans int
	j--
	if c > 0 {
		ans = buses[len(buses)-1]
	} else {
		ans = passengers[j] - 1
		j--
	}

	for ; j >= 0; j-- {
		if ans != passengers[j] {
			break
		}
		ans--
	}

	return ans
}
