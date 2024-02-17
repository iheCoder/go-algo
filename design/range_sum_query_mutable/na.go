package range_sum_query_mutable

type node struct {
	low, high   int
	val         int
	left, right *node
	boundary    int
}

func (n *node) Update(k, v int) int {
	if k < n.low || k > n.high {
		return 0
	}

	var d int
	if n.low == n.high {
		d = v - n.val
		n.val = v
		return d
	}

	if n.left != nil && n.boundary >= k {
		d = n.left.Update(k, v)
	} else if n.right != nil && n.boundary < k {
		d = n.right.Update(k, v)
	}
	n.val += d
	return d
}

func (n *node) SumRange(low, high int) int {
	if low > n.high || high < n.low {
		return 0
	}

	if n.low == low && n.high == high {
		return n.val
	}

	var lv, rv int
	if n.boundary >= low {
		lv = n.left.SumRange(low, minInt(n.boundary, high))
	}
	if high > n.boundary {
		rv = n.right.SumRange(maxInt(low, n.boundary+1), high)
	}
	return lv + rv
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

type tree struct {
	root *node
}

func (t *tree) Update(k, v int) {
	t.root.Update(k, v)
}

func (t *tree) SumRange(low, high int) int {
	return t.root.SumRange(low, high)
}

func NewSegmentTree(nums []int) *tree {
	var f func(low, high int) *node
	f = func(low, high int) *node {
		if low == high {
			return &node{
				low:      low,
				high:     high,
				val:      nums[low],
				boundary: low,
			}
		}

		b := low + (high-low)/2
		left, right := f(low, b), f(b+1, high)
		return &node{
			low:      low,
			high:     high,
			val:      left.val + right.val,
			left:     left,
			right:    right,
			boundary: b,
		}
	}
	return &tree{root: f(0, len(nums)-1)}
}

type NumArray struct {
	t *tree
}

func Constructor(nums []int) NumArray {
	return NumArray{t: NewSegmentTree(nums)}
}

func (this *NumArray) Update(index int, val int) {
	this.t.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.t.SumRange(left, right)
}

// 是不是可以平摊呢？采用前缀和+数组，建立更新缓冲区，当更新3次或者进行sumRange刷新缓冲区
// 可是这样更新prefixSum复杂度还是不变啊！感觉还不如不用前缀和
//const (
//	defaultRefreshFrequency = 3
//)
//
//type NumArray struct {
//	nums      []int
//	cache     map[int]int
//	prefixSum []int
//}
//
//func Constructor(nums []int) NumArray {
//	ps := make([]int, len(nums)+1)
//	for i, num := range nums {
//		ps[i+1] = ps[i] + num
//	}
//	return NumArray{
//		nums:      nums,
//		cache:     make(map[int]int),
//		prefixSum: ps,
//	}
//}
//
//func (this *NumArray) Update(index int, val int) {
//	this.cache[index] = val
//	if len(this.cache) > defaultRefreshFrequency {
//		this.refreshCache()
//	}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	if len(this.cache) > 0 {
//		this.refreshCache()
//	}
//	return this.prefixSum[right+1] - this.prefixSum[left]
//}
//
//func (this *NumArray) refreshCache() {
//	for i, v := range this.cache {
//		diff := v - this.nums[i]
//		for j := 1; j <= i+1; j++ {
//			this.prefixSum[j] += diff
//		}
//		this.nums[i] = v
//		delete(this.cache, i)
//	}
//}
