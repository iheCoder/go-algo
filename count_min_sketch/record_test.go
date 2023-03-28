package count_min_sketch

import (
	"fmt"
	"testing"
)

//
func TestSizeCal(t *testing.T) {
	x := 4
	n := (x-1)>>2 + 1
	fmt.Println(n)

	type test struct {
		num      int
		expected int
	}

	ts := []test{
		{
			1,
			1,
		},
		{15,
			1,
		},
		{16, 1},
		{17, 2},
	}

	for _, ti := range ts {
		r := (ti.num-1)>>4 + 1
		if r != ti.expected {
			t.Fatal("err")
		}
	}
}

func TestCMSRecord_Increase(t *testing.T) {
	// 给第18个添加数量3，并获取
	r := NewCMSCounter(20)
	r.Increase(17, 3)
	fmt.Println(r.GetCount(17))

	type test struct {
		index    int
		count    int
		expected int
	}

	ts := []test{
		{0, 1, 1},
		{15, 15, 15},
		{16, 1, 1},
		{16, 1, 2},
	}

	for _, ti := range ts {
		r.Increase(ti.index, ti.count)
		x := r.GetCount(ti.index)
		if ti.expected != x {
			t.Fatalf("bit record err: expected %d got %d", ti.expected, x)
		}
	}
	r.ResetByIndex(16)
	if r.GetCount(16) != 0 {
		t.Fatal("reset failed")
	}
}

func TestHexToDec(t *testing.T) {
	fmt.Printf("%d", delay2Mask)
}

func TestCMSRecord_Delay2(t *testing.T) {
	r := NewCMSCounter(20)
	r.Increase(17, 3)
	r.Increase(1, 8)
	r.Delay2()
	fmt.Println(r.GetCount(17))
	fmt.Println(r.GetCount(1))
}
