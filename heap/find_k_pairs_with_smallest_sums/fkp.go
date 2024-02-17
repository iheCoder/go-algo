package find_k_pairs_with_smallest_sums

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	ph := &ph{
		nums1: nums1,
		nums2: nums2,
		ps:    make([]*pair, 0),
	}

	for i := 0; i < k && i < m; i++ {
		heap.Push(ph, &pair{
			i: i,
			j: 0,
		})
	}

	var ans [][]int
	for len(ans) < k {
		x := heap.Pop(ph)
		xp := x.(*pair)
		ans = append(ans, []int{nums1[xp.i], nums2[xp.j]})
		if xp.j+1 < n {
			heap.Push(ph, &pair{
				i: xp.i,
				j: xp.j + 1,
			})
		}
	}

	return ans
}

type pair struct {
	i, j int
}

type ph struct {
	nums1, nums2 []int
	ps           []*pair
}

func (p *ph) Len() int {
	return len(p.ps)
}

func (p *ph) Less(i, j int) bool {
	sum1 := p.nums1[p.ps[i].i] + p.nums2[p.ps[i].j]
	sum2 := p.nums1[p.ps[j].i] + p.nums2[p.ps[j].j]
	return sum1 < sum2
}

func (p *ph) Swap(i, j int) {
	p.ps[i], p.ps[j] = p.ps[j], p.ps[i]
}

func (p *ph) Pop() interface{} {
	x := p.ps[len(p.ps)-1]
	p.ps = p.ps[:len(p.ps)-1]
	return x
}

func (p *ph) Push(i interface{}) {
	p.ps = append(p.ps, i.(*pair))
}

// 【1，2，3】，【2，3，4】无法应对这种情况
// 	var i, j int
//	var count int
//	var p1, p2 int
//	var ans [][]int
//	for count < k {
//		ans = append(ans, []int{nums1[i], nums2[j]})
//		count++
//		if count >= k {
//			break
//		}
//		if i+1 >= m {
//			j++
//			i = 0
//			continue
//		}
//		if j+1 >= n {
//			i++
//			j = 0
//			continue
//		}
//		p1, p2 = nums1[i+1]+nums2[j], nums1[i]+nums2[j+1]
//		if p1 < p2 {
//			i++
//		} else {
//			j++
//		}
//	}
//	return ans
