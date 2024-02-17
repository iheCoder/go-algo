package rec_area

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	w := minInt(ax2, bx2) - maxInt(ax1, bx1)
	h := minInt(ay2, by2) - maxInt(ay1, by1)
	var overlap int
	if w > 0 && h > 0 {
		overlap = w * h
	}
	total := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	return total - overlap
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
