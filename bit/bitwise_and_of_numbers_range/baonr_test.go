package bitwise_and_of_numbers_range

import (
	"fmt"
	"testing"
)

func TestBitLevel(t *testing.T) {
	x := 7
	fmt.Println(getBitLevel(x))
	x = 8
	fmt.Println(getBitLevel(x))
	x = 15
	fmt.Println(getBitLevel(x))

}
