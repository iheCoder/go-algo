package _sum_with_multiplicity

func threeSumMulti(arr []int, target int) int {
	const mod int = 1_000_000_007
	count := make([]int, 101)
	var uniqCount int
	for _, x := range arr {
		count[x]++
		if count[x] == 1 {
			uniqCount++
		}
	}

	var index int
	keys := make([]int, uniqCount)
	for i := 0; i < 101; i++ {
		if count[i] > 0 {
			keys[index] = i
			index++
		}
	}

	var ans int
	for i := 0; i < uniqCount; i++ {
		x := keys[i]
		t := target - x
		j, k := i, uniqCount-1

		for j <= k {
			y, z := keys[j], keys[k]
			if y+z < t {
				j++
			} else if y+z > t {
				k--
			} else {
				if i < j && j < k {
					ans += count[x] * count[y] * count[z]
				} else if i == j && k > j {
					ans += count[x] * (count[x] - 1) / 2 * count[z]
				} else if j == k && j > i {
					ans += count[x] * (count[z] - 1) * count[z] / 2
				} else {
					ans += count[x] * (count[x] - 1) * (count[x] - 2) / 6
				}
				ans %= mod

				j++
				k--
			}
		}
	}

	return ans
}

//func threeSumMulti(arr []int, target int) int {
//	const mod int = 10e9 + 7
//	sort.Ints(arr)
//	var ans int
//	n := len(arr)
//
//	for i := 0; i < n; i++ {
//		t := target - arr[i]
//		j, k := i+1, n-1
//		for j < k {
//			sum := arr[j] + arr[k]
//			if sum < t {
//				j++
//			} else if sum > t {
//				k--
//			} else if arr[j] != arr[k] {
//				l, r := 1, 1
//				for j+1 < k && arr[j] == arr[j+1] {
//					l++
//					j++
//				}
//				for k-1 > j && arr[k] == arr[k-1] {
//					r++
//					k--
//				}
//
//				ans += l * r
//				ans %= mod
//				j++
//				k--
//			} else {
//				ans += (k - j + 1) * (k - j) / 2
//				ans %= mod
//				break
//			}
//		}
//	}
//
//	return ans
//}

//func threeSumMulti(arr []int, target int) int {
//	m := make(map[int]int)
//	for _, num := range arr {
//		m[num]++
//	}
//
//	const mod int = 10e9 + 7
//
//	type pair struct {
//		num, count int
//	}
//	n := len(m)
//	ps := make([]*pair, 0, n)
//	for num, count := range m {
//		ps = append(ps, &pair{
//			num:   num,
//			count: count,
//		})
//	}
//	sort.Slice(ps, func(i, j int) bool {
//		return ps[i].num < ps[j].num
//	})
//
//	var ans int
//	for i := 0; i < n-2; i++ {
//		j, k := i, n-1
//		for j < k {
//			sum := ps[i].num + ps[j].num + ps[k].num
//			if sum > target {
//				k--
//			} else if sum < target {
//				j++
//			} else if i != j {
//				ans += ps[i].count * ps[j].count * ps[k].count
//				ans %= mod
//				j++
//			} else {
//				c := ps[i].count
//				ans += (c * (c - 1)) / 2 * ps[k].count
//				ans %= mod
//				j++
//			}
//		}
//	}
//
//	return ans
//}
