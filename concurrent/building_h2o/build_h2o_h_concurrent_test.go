package building_h2o

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestChannelCASMultiGoroutinesHC(t *testing.T) {
	count := 100
	hch := make(chan int, 2)
	och := make(chan int, 1)
	och <- 1
	var state int32

	for i := 0; i < count; i++ {
		go OChHC(hch, och, &state)
		go HChHC(hch, och, &state)
		go HChHC(hch, och, &state)
	}

	time.Sleep(3 * time.Second)
}

var hBlock = make(chan int, 1)

func HChHC(hch, och chan int, state *int32) {
	<-hch
	fmt.Println("H")

	atomic.AddInt32(state, -1)
	if atomic.LoadInt32(state) == 0 {
		och <- 1
	}
}

func OChHC(hch, och chan int, state *int32) {
	<-och
	fmt.Println("O")
	atomic.AddInt32(state, 2)
	hch <- 1
	hch <- 1
}

func TestPureChannelMultiGoroutinesHC(t *testing.T) {
	count := 100
	hch := make(chan int, 2)
	och0 := make(chan int, 1)
	och1 := make(chan int, 1)
	hch <- 1
	hch <- 1

	for i := 0; i < count; i++ {
		go OPureChHC(hch, och0, och1)
		go HPureChHC(hch, och0, och1)
		go HPureChHC(hch, och0, och1)
	}

	time.Sleep(3 * time.Second)
}

var hStateCh = make(chan int, 1)

func HPureChHC(hch, och0, och1 chan int) {
	<-hch
	fmt.Println("H")
	select {
	case hStateCh <- 1:
		och0 <- 1
	case <-hStateCh:
		och1 <- 1
	}
}

func OPureChHC(hch, och0, och1 chan int) {
	<-och0
	<-och1
	fmt.Println("O")
	hch <- 1
	hch <- 1
}
