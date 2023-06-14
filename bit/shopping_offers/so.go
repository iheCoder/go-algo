package shopping_offers

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	filterSpecial := make([][]int, 0, len(special))
	for _, s := range special {
		var normalPrice int
		for i, sc := range s[:n] {
			normalPrice += price[i] * sc
		}
		if s[n] < normalPrice {
			filterSpecial = append(filterSpecial, s)
		}
	}

	df := make(map[string]int)

	var dfs func(curNeeds []byte) int
	dfs = func(curNeeds []byte) int {
		if res, has := df[string(curNeeds)]; has {
			return res
		}

		// 求取当前need采用单价需要的价格
		var minPrice int
		for i, cn := range curNeeds {
			minPrice += price[i] * int(cn)
		}

		// 遍历所有的大礼包，回溯取最小值
		var invalid bool
		for _, fs := range filterSpecial {
			var nextNeeds []byte
			invalid = false
			for i, c := range fs[:n] {
				diff := int(curNeeds[i]) - c
				if diff < 0 {
					invalid = true
					break
				}
				nextNeeds = append(nextNeeds, byte(diff))
			}
			if !invalid {
				minPrice = minInt(minPrice, dfs(nextNeeds)+fs[n])
			}
		}
		df[string(curNeeds)] = minPrice
		return minPrice
	}
	curNeed := make([]byte, len(needs))
	for i, need := range needs {
		curNeed[i] = byte(need)
	}
	return dfs(curNeed)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 	// 过滤不需要计算的大礼包，只保留需要计算的大礼包
//	filterSpecial := [][]int{}
//	for _, s := range special {
//		totalCount, totalPrice := 0, 0
//		for i, c := range s[:n] {
//			totalCount += c
//			totalPrice += c * price[i]
//		}
//		if totalCount > 0 && totalPrice > s[n] {
//			filterSpecial = append(filterSpecial, s)
//		}
//	}
//
//	dp := make(map[string]int)
//	var dfs func(curNeeds []byte) int
//	dfs = func(curNeeds []byte) int {
//		if res, has := dp[string(curNeeds)]; has {
//			return res
//		}
//		// 原价价格
//		var minPrice int
//		for i, p := range price {
//			minPrice += int(curNeeds[i]) * p
//		}
//		nextNeeds := make([]byte, n)
//		var buyForbid bool
//		for _, s := range filterSpecial {
//			buyForbid = false
//			for i, need := range curNeeds {
//				if need < byte(s[i]) {
//					buyForbid = true
//					break
//				}
//				// 删除下次需要递归的need
//				nextNeeds[i] = need - byte(s[i])
//			}
//			// 只要不超出购买数量，就继续递归，求最小值
//			if !buyForbid {
//				minPrice = minInt(minPrice, dfs(nextNeeds)+s[n])
//			}
//		}
//		dp[string(curNeeds)] = minPrice
//		return minPrice
//	}
//	curNeeds := make([]byte, n)
//	for i, need := range needs {
//		curNeeds[i] = byte(need)
//	}
//	return dfs(curNeeds)
