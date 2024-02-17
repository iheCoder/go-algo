package print_zero_even_odd

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	n := 500

	zc := make(chan int, 1)
	ec := make(chan int, 1)
	oc := make(chan int, 1)
	zc <- 1
	for i := 0; i < n; i++ {
		go zeroCh(zc, ec, oc)
		go evenCh(zc, ec)
		go oddCh(zc, oc)
	}

	time.Sleep(time.Second)
}

var zeroInnerCh = make(chan int, 1)

func zeroCh(zeroCh, evenCh, oddCh chan int) {
	<-zeroCh
	fmt.Println("0")
	select {
	case zeroInnerCh <- 1:
		oddCh <- 1
	case <-zeroInnerCh:
		evenCh <- 1
	}
}

var oddNum = 1
var evenNum = 2

func oddCh(zeroCh, oddCh chan int) {
	<-oddCh
	fmt.Println(oddNum)
	oddNum += 2
	zeroCh <- 1
}

func evenCh(zeroCh, evenCh chan int) {
	<-evenCh
	fmt.Println(evenNum)
	evenNum += 2
	zeroCh <- 1
}

func TestThreeChannel(t *testing.T) {
	n := 500

	zc := make(chan int, 1)
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	zc <- 1
	zeroOneCh <- 1
	for i := 0; i < n; i++ {
		go zero3Ch(zc, c1, c2, c3)
		go one3Ch(zc, c1)
		go two3Ch(zc, c2)
		go three3Ch(zc, c3)
	}

	time.Sleep(time.Second)
}

var (
	zeroOneCh   = make(chan int, 1)
	zeroTwoCh   = make(chan int, 1)
	zeroThreeCh = make(chan int, 1)
)

func zero3Ch(zeroCh, oneCh, twoCh, threeCh chan int) {
	<-zeroCh
	fmt.Println("0")
	select {
	case <-zeroOneCh:
		oneCh <- 1
		zeroTwoCh <- 1
	case <-zeroTwoCh:
		twoCh <- 1
		zeroThreeCh <- 1
	case <-zeroThreeCh:
		threeCh <- 1
		zeroOneCh <- 1
	}
}

var (
	oneNum   = 1
	twoNum   = 2
	threeNum = 3
)

func one3Ch(zeroCh, oneCh chan int) {
	<-oneCh
	fmt.Println(oneNum)
	oneNum += 3
	zeroCh <- 1
}

func two3Ch(zeroCh, twoCh chan int) {
	<-twoCh
	fmt.Println(twoNum)
	twoNum += 3
	zeroCh <- 1
}

func three3Ch(zeroCh, threeCh chan int) {
	<-threeCh
	fmt.Println(threeNum)
	threeNum += 3
	zeroCh <- 1
}
