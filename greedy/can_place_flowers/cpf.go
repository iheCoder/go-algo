package can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	var c int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 == len(flowerbed) || flowerbed[i+1] == 0) {
			c++
			flowerbed[i] = 1
			if c >= n {
				return true
			}
		}
	}
	return false
}
