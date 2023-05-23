package friends_of_appropriate_ages

import (
	"sort"
)

func numFriendRequests(ages []int) int {
	if len(ages) <= 1 {
		return 0
	}

	sort.Ints(ages)

	var count, lowLimit int
	var step int
	var i, j int
	i = len(ages) - 1
	lowLimit = getUnderAge(ages[i])
	for j = i - 1; j >= 0 && ages[j] > lowLimit; j-- {
	}
	step = i - j - 1
	count += step

	for i = len(ages) - 2; i >= 0; i-- {
		if ages[i] == ages[i+1] {
			count += step
			continue
		}
		lowLimit = getUnderAge(ages[i])
		// 在相同的情况下，通过step求得的lastStop是应该按照最后的那个相同的index和step计算
		// 我还是想尝试回到记录lastStop，计算step
		if ages[i] <= lowLimit {
			continue
		}
		if j > i {
			j = i
		}
		for ; j >= 0 && ages[j] > lowLimit; j-- {
		}
		step = i - j - 1
		count += step
	}

	return count
}

func getUnderAge(age int) int {
	return age/2 + 7
}
