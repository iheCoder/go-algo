package longest_consecutive_seq

type unionFinder struct {
	m        map[int]int
	maxCount int
	counts   map[int]int
}

// 合并x，若存在x+1，则将x+1以及x+1 children纳入x下，若存在x-1，则将x以及x children纳入x-1下
func (u *unionFinder) union(x int) {
	if _, ok := u.m[x]; ok {
		return
	}
	u.m[x] = x
	u.counts[x] = 1
	if p, ok := u.m[x-1]; ok {
		u.m[x] = p
		u.counts[p]++
		u.maxCount = maxInt(u.maxCount, u.counts[p])
	}
	if p, ok := u.m[x+1]; ok {
		y := x + 1
		u.counts[u.m[x]] += u.counts[p]
		u.maxCount = maxInt(u.maxCount, u.counts[u.m[x]])
		for ok {
			u.m[y] = u.m[x]
			y++
			_, ok = u.m[y]
		}
	}
}

func (u *unionFinder) getMaxCount() int {
	return u.maxCount
}

func longestConsecutiveUF(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uf := &unionFinder{
		m:        make(map[int]int),
		maxCount: 1,
		counts:   make(map[int]int),
	}
	for _, num := range nums {
		uf.union(num)
	}
	return uf.getMaxCount()
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	ass := make(map[int]int)
	repeat := make(map[int]bool)
	ans := 1
	var x, c int
	for _, num := range nums {
		if repeat[num] {
			continue
		}
		repeat[num] = true
		x = num + 1
		c = 1
		for m[x] {
			xc, ok := ass[x]
			if ok {
				c += xc
				break
			}
			repeat[x] = true
			c++
			x++
		}
		ass[num] = c
		ans = maxInt(ans, c)
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 		x := 1
//		ass[num] = &x
//		if c, ok := ass[num+1]; ok {
//			x = *c + 1
//			ass[num] = &x
//			ass[num+1] = &x
//			ans = maxInt(ans, x)
//		}
//		if _, ok := ass[num-1]; ok {
//			*(ass[num-1])++
//			ass[num] = ass[num-1]
//			ans = maxInt(ans, *ass[num])
//		}
