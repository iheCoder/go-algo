package min_genetic_mutation

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumHighestOneDigit(t *testing.T) {
	num1 := 12345
	num2 := 500
	num3 := 9

	fmt.Println(num1, "has only highest digit?", hasOnlyHighestDigit(num1)) // Should print false
	fmt.Println(num2, "has only highest digit?", hasOnlyHighestDigit(num2)) // Should print true
	fmt.Println(num3, "has only highest digit?", hasOnlyHighestDigit(num3))
}

func hasOnlyHighestDigit(n int) bool {
	// Convert the number to a string
	strNum := strconv.Itoa(n)

	// Check if the first character is the only nonzero digit
	return len(strNum) > 0 && strNum[0] != '0' && len(strNum) == 1
}

func TestMM(t *testing.T) {
	sg := "AACCGGTT"
	eg := "AAACGGTA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}

	sg = "AAAAAAAA"
	eg = "AAGTAAAA"
	bank = []string{"CAAAAAAA", "CCAAAAAA", "CCATAAAA", "CCGTAAAA", "CAGTAAAA", "AAGTAAAA"}
	t.Log(minMutation(sg, eg, bank))
}

func TestOneChange(t *testing.T) {
	x := 11223344
	y := 1123341
	//x = 22141111
	//y = 21341111
	t.Log(checkOneChange(x, y))
}
