package single_number2

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var two, one int
	for _, x := range nums {
		one = one ^ x&^two
		two = two ^ x&^one
	}

	return one
}

// 设采用(two, one)来表示数当前二进制数所在位1出现的次数
// 00代表没有出现或者出现3次 01代表出现一次，也就是要找的值 10代表出现2次
// 那么最后出现一次的值就是one的或结果，因为仅有出现一次数最后one位才会有值
// x代表nums中的数值
// 	if two == 0 {
//		if x == 0 {
//			one = 0
//		} else {
//			one = 1
//		}
//	} else {
//		if x == 0 {
//			one = 0
//		} else {
//			one = 0
//		}
//	}
// 引入^简化如下
// 	if two == 0 {
//		one ^= x
//	} else {
//		one = 0
//	}
// 再次引入&简化如下
// 	one = one ^ x&^two
