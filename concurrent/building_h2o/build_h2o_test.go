package building_h2o

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// (1 信号量 o release two signal h
// (2 channel o send two to allow h consume
// (3 mutex state
// (4

func TestChannelIfLen(t *testing.T) {
	count := 100
	hch := make(chan int, 2)
	och := make(chan int, 1)
	och <- 1
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 2*count; i++ {
			HChIfLen(hch, och)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			OChIfLen(hch, och)
		}
		wg.Done()
	}()

	wg.Wait()
	//time.Sleep(3*time.Second)
}

func HChIfLen(hch, och chan int) {
	<-hch // a
	fmt.Println("H")
	// 近乎成功，但是死锁
	// 因为len不是加锁的，所以可能channel数据已经存在数据，但还是被判断为空
	// 既然deadlock那么要么是a、c死锁，要么是b、d死锁
	// 假设a、c死锁，那么说明hch与och都为空，那么可能是b处缺少到och的发送，
	// 那么可能是原本数据为空，但是判断是认为还存在数据，因此导致a、c死锁，但显然b前面就是a，应该不会是这种情况
	// 假设b、d死锁，那么说明och、hch都恰好缓存满，那么可能是b处原本数据不为空，但判断为空，因此往里面多发送了
	// (1 O consume, print "O"
	// (2 O produce
	// (3 H consume, and send O
	// (4 O produce
	// (5 H consume, and send O ,H block
	// (6 O consume, and still extra remind
	// (7 O produce two (there send 3 to H)
	// (8 H consume one, and send extra to O
	// (9 O block, H consume one
	// (10 H try produce, H block
	// 这么长的一串都没有意义，就是H和O的数量不匹配
	if len(hch) == 0 {
		och <- 1 // b
	}
}

func OChIfLen(hch, och chan int) {
	<-och // c
	fmt.Println("O")
	hch <- 1 // d
	hch <- 1
}

// 在多个goroutine下无法达到预期目标
func TestChannelIfLenMul(t *testing.T) {
	count := 100
	hch := make(chan int, 2)
	och := make(chan int, 1)
	och <- 1

	for i := 0; i < count; i++ {
		go OChIfLenMul(hch, och)
		go HChIfLenMul(hch, och)
		go HChIfLenMul(hch, och)
	}

	time.Sleep(3 * time.Second)
}

func HChIfLenMul(hch, och chan int) {
	<-hch
	fmt.Println("H")
	if len(hch) == 0 {
		och <- 1
	}
}

func OChIfLenMul(hch, och chan int) {
	<-och
	fmt.Println("O")
	hch <- 1
	hch <- 1
}

func TestChannelTwoPassForO(t *testing.T) {
	count := 100
	hch := make(chan int, 2)
	och := make(chan int, 2)
	hch <- 1
	hch <- 1

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 2*count; i++ {
			HChTwoPassForO(hch, och)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			OChTwoPassForO(hch, och)
		}
		wg.Done()
	}()

	wg.Wait()
}

func HChTwoPassForO(hch, och chan int) {
	<-hch
	fmt.Println("H")
	och <- 1
}

func OChTwoPassForO(hch, och chan int) {
	<-och
	<-och
	fmt.Println("O")
	hch <- 1
	hch <- 1
}

func TestChannel(t *testing.T) {
	count := 100
	hch := make(chan int, 1)
	och := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 2*count; i++ {
			HCh(hch, och)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			OCh(hch, och)
		}
		wg.Done()
	}()

	wg.Wait()
}

var lock = make(chan int, 1)
var hInnerCh = make(chan int, 1)

func HCh(hch, och chan int) {
	<-hch // a
	fmt.Println("H")

	select {
	case hInnerCh <- 1:
		hch <- 1
	case <-hInnerCh:
		och <- 1
	}

	//lock <- 1
	//if len(hch) == 0 {
	//	och <- 1 // b
	//}
	//<-lock
}

func OCh(hch, och chan int) {
	fmt.Println("O")

	//lock <- 1
	hch <- 1
	<-och
	//<-lock
}

func TestChannelMultiGoroutines(t *testing.T) {
	count := 100
	hch := make(chan int, 1)
	och := make(chan int)
	hch <- 1

	for i := 0; i < count; i++ {
		go OChMul(hch, och)
		go HChMul(hch, och)
		go HChMul(hch, och)
	}

	time.Sleep(3 * time.Second)
}

func HChMul(hch, och chan int) {
	<-hch
	fmt.Println("H")

	select {
	case hInnerCh <- 1:
		hch <- 1
	case <-hInnerCh:
		och <- 1
	}
}

func OChMul(hch, och chan int) {
	<-och
	fmt.Println("O")
	hch <- 1
}

func TestCASStateTransition(t *testing.T) {
	count := 100
	wg := &sync.WaitGroup{}
	wg.Add(2)

	v := &atomic.Value{}
	v.Store(0)
	go func() {
		for i := 0; i < 2*count; i++ {
			HCAS(v)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			OCAS(v)
		}
		wg.Done()
	}()

	wg.Wait()
}

func HCAS(v *atomic.Value) {
	for {
		for v.CompareAndSwap(2, 1) {
			fmt.Println("H")
			v.CompareAndSwap(1, 1)
			return
		}
		for v.CompareAndSwap(1, 1) {
			fmt.Println("H")
			v.CompareAndSwap(1, 0)
			return
		}
	}
}

func OCAS(v *atomic.Value) {
	for !v.CompareAndSwap(0, 3) {
	}
	fmt.Println("O")
	for !v.CompareAndSwap(3, 2) {
	}
}

func TestCASStateTransitionMul(t *testing.T) {
	count := 100

	v := &atomic.Value{}
	v.Store(0)

	for i := 0; i < count; i++ {
		go HCASMul(v)
		go HCASMul(v)
		go OCASMul(v)
	}

	time.Sleep(3 * time.Second)
}

func HCASMul(v *atomic.Value) {
	for {
		for v.CompareAndSwap(4, 3) {
			fmt.Println("H")
			v.CompareAndSwap(3, 2)
			return
		}
		for v.CompareAndSwap(2, 1) {
			fmt.Println("H")
			v.CompareAndSwap(1, 0)
			return
		}
	}
}

func OCASMul(v *atomic.Value) {
	for !v.CompareAndSwap(0, 5) {
	}
	fmt.Println("O")
	for !v.CompareAndSwap(5, 4) {
	}
}

func TestCond(t *testing.T) {
	count := 100
	wg := &sync.WaitGroup{}
	wg.Add(3)

	mu := &sync.Mutex{}
	hc := sync.NewCond(mu)
	oc := sync.NewCond(mu)
	go func() {
		for i := 0; i < count; i++ {
			HCond(hc, oc)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			HCond(hc, oc)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			OCond(hc, oc)
		}
		wg.Done()
	}()

	wg.Wait()
}

var condState int

func HCond(hc, oc *sync.Cond) {
	hc.L.Lock()
	for condState == 0 {
		hc.Wait()
	}
	fmt.Println("H")
	condState--
	oc.Broadcast()
	hc.L.Unlock()
}

func OCond(hc, oc *sync.Cond) {
	oc.L.Lock()
	for condState > 0 {
		oc.Wait()
	}
	fmt.Println("O")
	condState = 2
	hc.Signal()
	oc.L.Unlock()
}

func TestMutexMul(t *testing.T) {
	count := 100

	mu := &sync.Mutex{}
	for i := 0; i < count; i++ {
		go HMutex(mu)
		go HMutex(mu)
		go OMutex(mu)
	}

	time.Sleep(3 * time.Second)
}

var muState int

func HMutex(mu *sync.Mutex) {
	mu.Lock()
	for muState == 0 {
		mu.Unlock()
		mu.Lock()
	}
	fmt.Println("H")
	muState--
	mu.Unlock()
}

func OMutex(mu *sync.Mutex) {
	mu.Lock()
	for muState > 0 {
		mu.Unlock()
		mu.Lock()
	}
	fmt.Println("O")
	muState = 2
	mu.Unlock()
}

func TestSem(t *testing.T) {
	count := 100

	hsem := semaphore.NewWeighted(2)
	osem := semaphore.NewWeighted(2)
	osem.Acquire(context.Background(), 2)

	for i := 0; i < count; i++ {
		go HSem(hsem, osem)
		go HSem(hsem, osem)
		go OSem(hsem, osem)
	}

	time.Sleep(3 * time.Second)
}

func HSem(hsem, osem *semaphore.Weighted) {
	hsem.Acquire(context.Background(), 1)
	fmt.Println("H")
	osem.Release(1)
}

func OSem(hsem, osem *semaphore.Weighted) {
	osem.Acquire(context.Background(), 2)
	fmt.Println("O")
	hsem.Release(2)
}
