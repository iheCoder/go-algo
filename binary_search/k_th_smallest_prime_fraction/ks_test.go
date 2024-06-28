package k_th_smallest_prime_fraction

import "testing"

func TestKS(t *testing.T) {
	arr := []int{1, 2, 3, 5}
	k := 3
	t.Log(kthSmallestPrimeFraction(arr, k))
}
