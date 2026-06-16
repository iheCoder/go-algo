package cinema_seat_allocation

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	occupied := make(map[int]int)
	for _, seat := range reservedSeats {
		row, col := seat[0], seat[1]
		if col == 1 || col == 10 {
			continue
		}

		occupied[row] |= 1 << (col - 2)
	}

	ans := (n - len(occupied)) * 2
	for _, oc := range occupied {
		cnt := 0
		if oc&0b11110000 == 0 {
			cnt++
		}
		if oc&0b00001111 == 0 {
			cnt++
		}
		if cnt == 0 && oc&0b00111100 == 0 {
			cnt = 1
		}
		ans += cnt
	}

	return ans
}
