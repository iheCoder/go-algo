package num_of_recent_calls

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{q: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.q) > 0 {
		if t-this.q[0] <= 3000 {
			break
		}
		this.q = this.q[1:]
	}
	this.q = append(this.q, t)
	return len(this.q)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
