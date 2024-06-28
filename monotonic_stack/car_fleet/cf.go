package car_fleet

import "sort"

func carFleet(target int, position []int, speed []int) int {
	type car struct {
		pos        int
		arriveTime float64
	}

	n := len(position)
	if n == 0 {
		return 0
	}
	cars := make([]*car, 0, n)
	for i := 0; i < n; i++ {
		cars = append(cars, &car{
			pos:        position[i],
			arriveTime: float64(target-position[i]) / float64(speed[i]),
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})

	var ans int
	for t := n - 1; t > 0; t-- {
		if cars[t-1].arriveTime > cars[t].arriveTime {
			ans++
		} else {
			// 由于速度追到之后会均等，所以就直接到达时间等于前一个
			cars[t-1] = cars[t]
		}
	}
	// 最后一个独立车队和所有都在一个车队都要ans++
	ans++
	return ans
}
