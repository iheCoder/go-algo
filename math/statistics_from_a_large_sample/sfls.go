package statistics_from_a_large_sample

func sampleStats(count []int) []float64 {
	n := len(count)
	minI := 0
	for minI < n {
		if count[minI] > 0 {
			break
		}
		minI++
	}

	maxI := n - 1
	for maxI >= 0 {
		if count[maxI] > 0 {
			break
		}
		maxI--
	}

	var mc, mode int
	var sum int64
	var sumCount int
	for i, c := range count {
		if c > mc {
			mc = c
			mode = i
		}
		sum += int64(c * i)
		sumCount += c
	}
	mean := float64(sum) / float64(sumCount)
	mi := (sumCount + 1) / 2

	var sumHalf int
	var median float64
	for i, c := range count {
		sumHalf += c
		if sumHalf >= mi {
			if sumCount%2 == 1 || sumHalf != mi {
				median = float64(i)
			} else {
				j := i + 1
				for j < n {
					if count[j] > 0 {
						median = float64(i+j) / 2
						break
					}
					j++
				}
			}
			break
		}
	}

	return []float64{
		float64(minI),
		float64(maxI),
		mean,
		median,
		float64(mode),
	}
}
