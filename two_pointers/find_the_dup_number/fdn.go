package find_the_dup_number

import "fmt"

func findDuplicateFac(nums []int) int {
	r := factorialInt(len(nums) - 1)
	for i := 0; i < len(nums); i++ {
		if r%nums[i] > 0 {
			return nums[i]
		}
		r = r / nums[i]
	}
	return 1
}

func factorialInt(n int) int {
	r := 1
	for i := 1; i < n+1; i++ {
		r *= i
	}
	return r
}

func findDuplicate(nums []int) int {
	var slow, fast int
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fmt.Printf("nums %v slow %d\n", nums, slow)
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
