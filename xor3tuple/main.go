package xor3tuple

func main() {

}

func countTriplets(arr []int) int {
	total := 0
	length := len(arr)

	for i := 0; i < length-1; i++ {
		xor := arr[i]
		for j := i + 1; j < length; j++ {
			xor ^= arr[j]
			if xor == 0 {
				total += j - i
			}
		}
	}

	return total
}
