package shuffle_an_array

import "math/rand"

type Solution struct {
	oa, ca []int
}

func Constructor(nums []int) Solution {
	ca := make([]int, len(nums))
	copy(ca, nums)
	return Solution{
		oa: nums,
		ca: ca,
	}
}

func (this *Solution) Reset() []int {
	return this.oa
}

func (this *Solution) Shuffle() []int {
	n := len(this.ca)
	for i := 0; i < n; i++ {
		j := rand.Int63n(int64(n))
		this.ca[i], this.ca[j] = this.ca[j], this.ca[i]
	}
	return this.ca
}
