package find_median_from_data_stream

import "github.com/emirpasic/gods/trees/redblacktree"

type MedianFinder struct {
	nums *redblacktree.Tree
}

func Constructor() MedianFinder {
	return MedianFinder{nums: redblacktree.NewWithIntComparator()}
}

func (this *MedianFinder) AddNum(num int) {

}

func (this *MedianFinder) FindMedian() float64 {
	return 0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
