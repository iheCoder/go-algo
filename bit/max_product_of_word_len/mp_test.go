package max_product_of_word_len

import "testing"

func TestMP(t *testing.T) {
	words := []string{"abcde", "ahjks", "x", "ab", "cd"}
	t.Log(maxProduct(words))
}
