package range_sum_query_mutable

import "testing"

func TestNA(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	na := Constructor(nums)
	//t.Log(na.SumRange(0,2))
	//na.Update(2,4)
	//t.Log(na.SumRange(0,2))
	// <=不太对
	//t.Log(na.SumRange(1,3))
	//na.Update(2,5)
	//t.Log(na.SumRange(1,3))

	t.Log(na.SumRange(1, 1))
}
