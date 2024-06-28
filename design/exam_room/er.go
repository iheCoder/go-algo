package exam_room

import (
	"container/heap"
	"sort"
)

type ExamRoom struct {
	n     int
	seats *SortedSet
	q     priorityQueue
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		n:     n,
		seats: NewSortedSet(),
		q:     make(priorityQueue, 0),
	}
}

func (er *ExamRoom) Seat() int {
	// 若无人做，则取0位
	if er.seats.Len() == 0 {
		er.seats.Add(0)
		return 0
	}

	// 取最左边、最右边空座位区间
	left, right := er.seats.First(), er.n-1-er.seats.Last()

	for er.seats.Len() >= 2 {
		p := er.q[0]
		if er.seats.Contains(p[0]) && er.seats.Contains(p[1]) && er.seats.higher(p[0]) == p[1] {
			d := p[1] - p[0]
			if d/2 < right || d/2 <= left {
				break
			}

			heap.Pop(&er.q)
			heap.Push(&er.q, []int{p[0], p[0] + d/2})
			heap.Push(&er.q, []int{p[0] + d/2, p[1]})
			er.seats.Add(p[0] + d/2)
			return p[0] + d/2
		}
		heap.Pop(&er.q)
	}

	if right > left {
		heap.Push(&er.q, []int{er.seats.Last(), er.n - 1})
		er.seats.Add(er.n - 1)
		return er.n - 1
	} else {
		heap.Push(&er.q, []int{0, er.seats.First()})
		er.seats.Add(0)
		return 0
	}
}

func (er *ExamRoom) Leave(p int) {
	if p != er.seats.First() && p != er.seats.Last() {
		index := er.seats.indexOf(p)
		prev, next := er.seats.set[index-1], er.seats.set[index+1]
		heap.Push(&er.q, []int{prev, next})
	}
	er.seats.Remove(p)
}

type priorityQueue [][]int

func (q priorityQueue) Len() int {
	return len(q)
}

func (q priorityQueue) Less(i, j int) bool {
	d1, d2 := q[i][1]-q[i][0], q[j][1]-q[j][0]
	if d1/2 < d2/2 || (d1/2 == d2/2 && q[i][0] > q[j][0]) {
		return false
	}
	return true
}

func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *priorityQueue) Push(x interface{}) {
	item := x.([]int)
	*q = append(*q, item)
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}

type SortedSet struct {
	set []int
}

func NewSortedSet() *SortedSet {
	return &SortedSet{
		set: make([]int, 0),
	}
}

func (s *SortedSet) Add(val int) {
	if !s.Contains(val) {
		s.set = append(s.set, val)
		sort.Ints(s.set)
	}
}

func (s *SortedSet) Remove(val int) {
	index := s.indexOf(val)
	if index != len(s.set) {
		s.set = append(s.set[:index], s.set[index+1:]...)
	}
}

func (s *SortedSet) Contains(val int) bool {
	if len(s.set) == 0 {
		return false
	}

	index := s.indexOf(val)
	if index == len(s.set) || s.set[index] != val {
		return false
	}
	return true
}

func (s *SortedSet) indexOf(val int) int {
	return sort.SearchInts(s.set, val)
}

func (s *SortedSet) Len() int {
	return len(s.set)
}

func (s *SortedSet) First() int {
	return s.set[0]
}

func (s *SortedSet) Last() int {
	return s.set[len(s.set)-1]
}

func (s *SortedSet) higher(target int) int {
	for _, num := range s.set {
		if num > target {
			return num
		}
	}
	return -1
}
