package my_calendar

type MyCalendar struct {
	// lazy直接记录当前区间是不是存在
	// trees仅记录是否存在子区间
	subintervalExists, lazy map[int]bool
}

func Constructor() MyCalendar {
	return MyCalendar{
		subintervalExists: make(map[int]bool),
		lazy:              make(map[int]bool),
	}
}

func (mc *MyCalendar) query(start, end, l, r, idx int) bool {
	if start > r || end < l {
		return false
	}

	if mc.lazy[idx] {
		return true
	}

	if l >= start && r <= end {
		return mc.subintervalExists[idx]
	}

	mid := l + (r-l)/2
	return mc.query(start, end, l, mid, 2*idx) || mc.query(start, end, mid+1, r, 2*idx+1)
}

func (mc *MyCalendar) update(start, end, l, r, idx int) {
	if start > r || end < l {
		return
	}

	mc.subintervalExists[idx] = true

	if start <= l && end >= r {
		mc.lazy[idx] = true
		return
	}

	mid := l + (r-l)/2
	mc.update(start, end, l, mid, idx*2)
	mc.update(start, end, mid+1, r, idx*2+1)
	if mc.lazy[idx*2] && mc.lazy[idx*2+1] {
		mc.lazy[idx] = true
	}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	if mc.query(start, end-1, 0, 1e9, 1) {
		return false
	}
	mc.update(start, end-1, 0, 1e9, 1)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
