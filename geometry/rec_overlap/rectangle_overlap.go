package rec_overlap

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	var w, h int
	if (rec1[0]-rec2[2])*(rec1[2]-rec2[0]) < 0 {
		w = 1
	}
	if (rec1[1]-rec2[3])*(rec1[3]-rec2[1]) < 0 {
		h = 1
	}
	return w*h > 0
}
