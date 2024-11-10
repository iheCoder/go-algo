package distant_barcodes

func rearrangeBarcodes(barcodes []int) []int {
	if len(barcodes) < 2 {
		return barcodes
	}

	counts := make(map[int]int)
	for _, b := range barcodes {
		counts[b]++
	}

	evenIndex, oddIndex := 0, 1
	halfLen := len(barcodes) / 2
	res := make([]int, len(barcodes))
	for b, c := range counts {
		for c > 0 && c <= halfLen && oddIndex < len(barcodes) {
			res[oddIndex] = b
			c--
			oddIndex += 2
		}
		for c > 0 {
			res[evenIndex] = b
			c--
			evenIndex += 2
		}
	}

	return res
}
