package print_in_order

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// 1. init ac, bc channel
// 2. printA read ac,write bc, printB read bc, write ac
func TestChannelAB(t *testing.T) {
	ac := make(chan int, 1)
	bc := make(chan int, 1)
	ac <- 1
	count := 100

	go func() {
		for i := 0; i < count; i++ {
			printA(ac, bc)
		}
	}()
	go func() {
		for i := 0; i < count; i++ {
			printB(ac, bc)
		}
	}()
	time.Sleep(time.Second)
}

func printA(ac chan int, bc chan int) {
	<-ac
	fmt.Println("a")
	bc <- 1
}

func printB(ac chan int, bc chan int) {
	<-bc
	fmt.Println("b")
	ac <- 1

}

func TestFetchFromChannelThatNotPut(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		select {
		case <-ch:
			t.Log("SHOULD AT LAST")
		}
	}()
	time.Sleep(time.Second)
	t.Log("------")
	ch <- 1
	time.Sleep(time.Second)
}

func TestOneChannel(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		printAOC(ch)
	}()
	go func() {
		printBOC(ch)
	}()

	time.Sleep(time.Second)
}

func TestTryOneChannel(t *testing.T) {
	ch := make(chan int, 1)
	count := 100
	go func() {
		for i := 0; i < count; i++ {
			printAOC(ch)
		}
	}()

	go func() {
		for i := 0; i < count; i++ {
			printBOC(ch)
		}
	}()

	time.Sleep(time.Second)
}

func TestCAS(t *testing.T) {
	count := 100
	v := &atomic.Value{}
	v.Store(0)
	go func() {
		for i := 0; i < count; i++ {
			printACAS(v)
		}
	}()

	go func() {
		for i := 0; i < count; i++ {
			printBCAS(v)
		}
	}()

	time.Sleep(time.Second)
}

func printACAS(v *atomic.Value) {
	for !v.CompareAndSwap(0, 1) {
	}
	fmt.Println("a")
	for !v.CompareAndSwap(1, 2) {
	}
}

func printBCAS(v *atomic.Value) {
	for !v.CompareAndSwap(2, 1) {
	}
	fmt.Println("b")
	for !v.CompareAndSwap(1, 0) {
	}
}

func printATryOC(ch chan int) {
	<-ch
	fmt.Println("a")
}

func printBTryOC(ch chan int) {
	ch <- 1
	fmt.Println("b")
}

// 我能理解不按照顺序打印，可是我不能理解为什么不是一个接着一个呢？
func printAOC(ch chan int) {
	for i := 0; i < 100; i++ {
		select {
		case <-ch:
			fmt.Println("a")
		}
	}
}

func printBOC(ch chan int) {
	for i := 0; i < 100; i++ {
		select {
		case ch <- 1:
			fmt.Println("b")
		}
	}
}
