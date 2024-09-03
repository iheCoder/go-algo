package robot_bounded_in_circle

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	dx, dy := 0, 1
	for _, c := range instructions {
		if c == 'G' {
			x += dx
			y += dy
		} else if c == 'L' {
			dx, dy = -dy, dx
		} else {
			dx, dy = dy, -dx
		}
	}

	return x == 0 && y == 0 || dx != 0 || dy != 1
}

// 	var x, y int
//	dx, dy := 0, 1
//	for _, ins := range instructions {
//		switch ins {
//		case 'G':
//			x += dx
//			y += dy
//		case 'L':
//			dx = -dy
//			dy = dx
//		case 'R':
//			dx = dy
//			dy = -dx
//		}
//	}
//
//	return (x == 0 && y == 0) || !(dx == 0 && dy == 1)
