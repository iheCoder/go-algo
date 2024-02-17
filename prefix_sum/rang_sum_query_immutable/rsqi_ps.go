package rang_sum_query_immutable

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	sums := make([]int, n+1)
	sums[0] = 0
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}
