package top_k_frequent_elements

import (
	"math/rand"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	arr := make([]int, 0, len(m))
	for num, _ := range m {
		arr = append(arr, num)
	}
	rand.Seed(time.Now().Unix())

	if len(arr) <= k {
		return arr
	}

	var f func(start, end, n int)
	f = func(start, end, n int) {
		if n <= 0 {
			return
		}

		//arr[start], arr[end] = arr[end], arr[start]
		xi := rand.Intn(end-start+1) + start
		arr[xi], arr[end] = arr[end], arr[xi]
		x := m[arr[end]]
		i, j := start, end-1
		for {
			for i < end && m[arr[i]] > x {
				i++
			}
			for j >= 0 && m[arr[j]] <= x {
				j--
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i], arr[end] = arr[end], arr[i]
		l := i - start
		if l == n {
			return
		}
		if l > n {
			f(start, i-1, n)
		} else {
			f(i, end, n-l)
		}
	}

	f(0, len(arr)-1, k)
	return arr[:k]
}
