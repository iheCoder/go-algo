package barber_problem

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	rand.Seed(time.Now().Unix())
	var st time.Duration
	go barberCond(cond)

	count := 100
	for i := 0; i < count; i++ {
		go customerCond(cond)
		st = time.Duration(rand.Intn(100))
		time.Sleep(st * time.Nanosecond)
	}

	time.Sleep(3 * time.Second)
}

var waits int

func barberCond(cond *sync.Cond) {
	for {
		cond.L.Lock()
		for waits == 0 {
			cond.Wait()
		}
		fmt.Println("haircutting...")
		time.Sleep(time.Microsecond)
		fmt.Println("complete haircut")
		waits--
		cond.L.Unlock()
	}
}

func customerCond(cond *sync.Cond) {
	cond.L.Lock()
	waits++
	if waits == 1 {
		cond.Signal()
	}
	cond.L.Unlock()
}

func TestChannel(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var st time.Duration
	ch := make(chan int)
	go barberCh(ch)

	count := 100
	for i := 0; i < count; i++ {
		go consumerCh(ch)
		st = time.Duration(rand.Intn(100))
		time.Sleep(st * time.Nanosecond)
	}

	time.Sleep(3 * time.Second)
}

func barberCh(ch chan int) {
	for {
		<-ch
		fmt.Println("haircutting...")
		time.Sleep(time.Microsecond)
		fmt.Println("complete haircut")
	}
}

func consumerCh(ch chan int) {
	ch <- 1
}
