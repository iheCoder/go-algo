package valid_square

import "fmt"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if p1[0] == p2[0] && p3[0] == p4[0] && p1[0] == p3[0] && p1[1] == p2[1] && p3[1] == p4[1] && p1[1] == p3[1] {
		return false
	}

	cx := p1[0] + p2[0] + p3[0] + p4[0]
	cy := p1[1] + p2[1] + p3[1] + p4[1]
	m := make(map[string]bool)

	p1 = moveTo0(p1, cx, cy)
	m[getKey(p1)] = true
	p2 = moveTo0(p2, cx, cy)
	m[getKey(p2)] = true
	p3 = moveTo0(p3, cx, cy)
	m[getKey(p3)] = true
	p4 = moveTo0(p4, cx, cy)
	m[getKey(p4)] = true

	fp1 := flip(p1)
	if !m[getKey(fp1)] {
		return false
	}
	fp2 := flip(p2)
	if !m[getKey(fp2)] {
		return false
	}
	fp3 := flip(p3)
	if !m[getKey(fp3)] {
		return false
	}
	fp4 := flip(p4)
	if !m[getKey(fp4)] {
		return false
	}
	return true
}

func flip(p []int) []int {
	return []int{p[1], -p[0]}
}

func moveTo0(p []int, cx, cy int) []int {
	return []int{4*p[0] - cx, 4*p[1] - cy}
}

func getKey(p []int) string {
	return fmt.Sprintf("%d_%d", p[0], p[1])
}

//func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
//	s1, s2, s3 := calculateSideLen(p1, p2), calculateSideLen(p1, p3), calculateSideLen(p1, p4)
//	var mxl, mil int
//	if s1 > s2 && s2 == s3 {
//		mxl = s1
//		mil = s2
//		if checkInOneLine(p1, p3, p4) {
//			return false
//		}
//		if !checkCos0(p1, p3, p2, p4) {
//			return false
//		}
//	} else if s2 > s1 && s1 == s3 {
//		mxl = s2
//		mil = s1
//		if checkInOneLine(p1, p2, p4) {
//			return false
//		}
//		if !checkCos0(p1, p2, p3, p4) {
//			return false
//		}
//	} else if s3 > s1 && s1 == s2 {
//		mxl = s3
//		mil = s1
//		if checkInOneLine(p1, p2, p3) {
//			return false
//		}
//		if !checkCos0(p1, p2, p3, p4) {
//			return false
//		}
//	} else {
//		return false
//	}
//
//	return mxl == 2*mil
//}
//
//func calculateSideLen(p1, p2 []int) int {
//	xd := p1[0] - p2[0]
//	yd := p1[1] - p2[1]
//	return xd*xd + yd*yd
//}
//
//func checkInOneLine(p1, p2, p3 []int) bool {
//	s1, s2 := p2[0]-p1[0], p2[1]-p1[1]
//	s3, s4 := p3[0]-p1[0], p3[1]-p1[1]
//	return s1*s4 == s2*s3
//}
//
//func checkCos0(p1, p2, p3, p4 []int) bool {
//	v1 := []int{p1[0] - p2[0], p1[1] - p2[1]}
//	v2 := []int{p3[0] - p4[0], p3[1] - p4[1]}
//	x := v1[0]*v2[0] + v1[1]*v2[1]
//	return x == 0
//}
