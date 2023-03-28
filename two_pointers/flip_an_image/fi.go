package flip_an_image

func flipAndInvertImage(image [][]int) [][]int {
	if len(image) == 0 {
		return image
	}

	n := len(image)
	for i := 0; i < n; i++ {
		j, k := 0, n-1
		for j <= k {
			image[i][j], image[i][k] = 1^image[i][k], 1^image[i][j]
			j++
			k--
		}
	}

	return image
}
