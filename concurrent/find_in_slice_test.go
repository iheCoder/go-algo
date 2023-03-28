package concurrent

import (
	"fmt"
	"testing"
	"time"
)

// 关闭可以利用select channel

type goFind struct {
	done  chan bool
	timer *time.Timer
}

func (f *goFind) findInSlice(x int, items []int) {
	for _, item := range items {
		if x == item {
			// 若前一个goroutine已经找到，那么这个是否会重复found it，甚至卡死呢，当然会的
			fmt.Println("Found it!")
			close(f.done)
			return
		}
		select {
		// 无法达到给一个，多个就能获取done结束
		case <-f.done:
			return
		case <-f.timer.C:
			fmt.Println("Timeout! Not Found")
			close(f.done)
		default:
		}
	}
}

func TestGoFind(t *testing.T) {
	items := []int{1, 5, 3, 56, 2134, 124, 354, 0, 45, 2, 44}
	gf := &goFind{
		done:  make(chan bool, 1),
		timer: time.NewTimer(5 * time.Millisecond),
	}
	go gf.findInSlice(44, items[:3])
	go gf.findInSlice(44, items[3:6])
	go gf.findInSlice(44, items[6:9])
	go gf.findInSlice(44, items[9:])
	time.Sleep(10 * time.Second)
}

func TestReceiveWithoutReceiver(t *testing.T) {
	ch := make(chan int, 1)
	go receiveWithoutReceiver(ch)
	go receiveWithoutReceiver(ch)
	go receiveWithoutReceiver(ch)
	go func() {
		ch <- 1
	}()
	go receiveWithoutReceiver(ch)
	go receiveWithoutReceiver(ch)
	time.Sleep(time.Second)
}

func receiveWithoutReceiver(ch chan int) {
	select {
	case <-ch:
		fmt.Println("a lot")
	}
}
