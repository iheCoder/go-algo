package bitwise_and_of_numbers_range

import (
	"fmt"
	"math/bits"
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

func getBitLevel(x int) int {
	if x == 0 {
		return 0
	}
	l := 0
	for x != 0 {
		x >>= 1
		l++
	}
	return l
}

func TestRangeBitwiseAnd(t *testing.T) {
	type testData struct {
		left, right int
		expected    int
	}
	tds := []testData{
		{5, 7, 4},
	}

	for _, td := range tds {
		got := rangeBitwiseAnd(td.left, td.right)
		if got != td.expected {
			t.Fatalf("left %d right %d got %d but expect %d", td.left, td.right, got, td.expected)
		}
	}
}

func TestBitOp(t *testing.T) {
	x := 1
	t.Log(x & -x)
	x = 2
	t.Log(x & -x)
	x = 3
	t.Log(x & -x)
	x = 4
	t.Log(x & -x)
	x = 5
	t.Log(x & -x)
}

func TestHighest1Bit(t *testing.T) {
	for i := 1; i < 1<<31-1; i++ {
		got := getHighest1Bit(i)
		expect := bits.Len(uint(i))
		if got+1 != expect {
			t.Fatalf("%d got %d but expect %d", i, got, expect)
		}
	}
}
