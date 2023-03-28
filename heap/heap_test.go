package heap

import (
	"fmt"
	"testing"
)

func TestHeapPop(t *testing.T) {
	i1 := &Item{
		Value:    1,
		Priority: 1,
		Index:    0,
	}
	i2 := &Item{
		Value:    2,
		Priority: 2,
		Index:    1,
	}
	i3 := &Item{
		Value:    3,
		Priority: 3,
		Index:    2,
	}

	pq := &PriorityQueue{
		i3,
		i2,
		i1,
	}

	x := Pop(pq)
	xi := x.(*Item)
	fmt.Println(xi.Priority)

	pq1 := &PriorityQueue{
		i3,
		i2,
	}
	x1 := Pop(pq1)
	x1i := x1.(*Item)
	fmt.Println(x1i.Priority)

	// what if when pop empty
	pq2 := &PriorityQueue{}
	x2 := Pop(pq2)
	x2i := x2.(*Item)
	fmt.Println(x2i.Priority)
}
