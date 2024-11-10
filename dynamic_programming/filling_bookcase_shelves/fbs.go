package filling_bookcase_shelves

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		var maxHeight, curWidth int
		for j := i; j >= 0; j-- {
			curWidth += books[j][0]
			if curWidth > shelfWidth {
				break
			}

			maxHeight = maxInt(maxHeight, books[j][1])
			dp[i+1] = minInt(dp[i+1], dp[j]+maxHeight)
		}
	}

	return dp[n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
