package snapshot_array

import "sort"

type SnapshotArray struct {
	snapVersion int
	data        [][][2]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapVersion: 0,
		data:        make([][][2]int, length),
	}
}

func (s *SnapshotArray) Set(index int, val int) {
	s.data[index] = append(s.data[index], [2]int{s.snapVersion, val})
}

func (s *SnapshotArray) Snap() int {
	s.snapVersion++
	return s.snapVersion - 1
}

func (s *SnapshotArray) Get(index int, snap_id int) int {
	y := s.data[index]
	x := sort.Search(len(y), func(i int) bool {
		return y[i][0] > snap_id
	})
	if x == 0 {
		return 0
	}

	return y[x-1][1]
}

// over mem
// type SnapshotArray struct {
//	snapValues [][]int
//	length     int
//	snapID     int
//}
//
//func Constructor(length int) SnapshotArray {
//	sv := make([][]int, length)
//	for i := 0; i < length; i++ {
//		sv[i] = append(sv[i], 0)
//	}
//
//	return SnapshotArray{
//		snapValues: sv,
//		length:     length,
//	}
//}
//
//func (s *SnapshotArray) Set(index int, val int) {
//	n := len(s.snapValues[index])
//	s.snapValues[index][n-1] = val
//}
//
//func (s *SnapshotArray) Snap() int {
//	for i := 0; i < s.length; i++ {
//		last := s.snapValues[i][len(s.snapValues[i])-1]
//		s.snapValues[i] = append(s.snapValues[i], last)
//	}
//
//	defer func() { s.snapID++ }()
//	return s.snapID
//}
//
//func (s *SnapshotArray) Get(index int, snap_id int) int {
//	return s.snapValues[index][snap_id]
//}
