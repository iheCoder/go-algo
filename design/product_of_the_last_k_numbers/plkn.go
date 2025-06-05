package product_of_the_last_k_numbers

type ProductOfNumbers struct {
	// 存储所有加入的数字
	nums []int
	// 记录前缀中 0 的数量（用于快速判断区间是否包含 0）
	zero []int
	// 记录每个数字前一个非 0、非 1 数字的索引
	pre []int
	// 当前已经添加数字的总个数
	n int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		nums: make([]int, 0),
		zero: make([]int, 0),
		pre:  make([]int, 0),
		n:    0,
	}
}

func (pn *ProductOfNumbers) Add(num int) {
	pn.n++
	pn.nums = append(pn.nums, num)
	pn.pre = append(pn.pre, -1)

	if pn.n > 1 {
		if pn.nums[pn.n-2] != 1 && pn.nums[pn.n-2] != 0 {
			pn.pre[pn.n-1] = pn.n - 2
		} else {
			pn.pre[pn.n-1] = pn.pre[pn.n-2]
		}
	}

	if num == 0 {
		pn.zero = append(pn.zero, 1)
	} else {
		pn.zero = append(pn.zero, 0)
	}

	if pn.n > 1 {
		pn.zero[pn.n-1] += pn.zero[pn.n-2]
	}
}

func (pn *ProductOfNumbers) GetProduct(k int) int {
	zc := pn.zero[pn.n-1]
	if pn.n-k > 0 {
		zc -= pn.zero[pn.n-1-k]
	}

	if zc > 0 {
		return 0
	}

	ans := 1
	for i := pn.n - 1; i >= pn.n-k; {
		ans *= pn.nums[i]
		i = pn.pre[i]
	}

	return ans
}
