package loop_array

import "testing"

func TestLoopArray(t *testing.T) {
	// 若总共0、1、2
	// 删除index之前的元素
	// 若index为2，那么删除之后为1

	// 删除index指向的元素
	// 若index为1，则删除之后为1

	// 删除index之后的元素
	// 若index为0，删除之后为0

	// 删除index指向的元素且该元素为最后一个
	// 若index为2，删除之后为0

	// 删除index之后的元素，后面的元素是最后一个
	// 若index为1，删除之后为1

	// 删除index之前的元素，之前的元素是第一个
	// 若index为1，删除之后为1

	// 总之就是被删的不是index指向元素就一直指着实际的那个元素，若是指向的元素就到下一个

	type ts struct {
		index      int
		deli       int
		afterIndex int
	}

	tss := []ts{
		{
			index:      2,
			deli:       1,
			afterIndex: 1,
		},
		{
			index:      1,
			deli:       0,
			afterIndex: 0,
		},

		{
			index:      0,
			deli:       1,
			afterIndex: 0,
		},
		{
			index:      1,
			deli:       2,
			afterIndex: 1,
		},

		{
			index:      1,
			deli:       1,
			afterIndex: 1,
		},
		{
			index:      2,
			deli:       2,
			afterIndex: 0,
		},
		{
			index:      0,
			deli:       0,
			afterIndex: 0,
		},
	}

	for _, tsi := range tss {

	}
}
