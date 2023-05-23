package building_h2o

import (
	"fmt"
	"testing"
	"time"
)

func TestGPTQueue(t *testing.T) {
	count := 100

	h2o := NewH2O()
	for i := 0; i < count; i++ {
		go h2o.hydrogen(releaseHydrogen)
		go h2o.hydrogen(releaseHydrogen)
		go h2o.oxygen(releaseOxygen)
	}

	time.Sleep(3 * time.Second)
}

func releaseHydrogen() {
	fmt.Println("H")
}

func releaseOxygen() {
	fmt.Println("O")
}

type H2O struct {
	hQueue  chan int
	oQueue  chan int
	waiting int
}

func NewH2O() *H2O {
	return &H2O{
		hQueue: make(chan int, 2),
		oQueue: make(chan int, 1),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.hQueue <- 1
	h2o.waiting++
	h2o.tryBond()
	<-h2o.oQueue
	releaseHydrogen()
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.oQueue <- 1
	h2o.waiting++
	h2o.tryBond()
	<-h2o.hQueue
	<-h2o.hQueue
	releaseOxygen()
}

func (h2o *H2O) tryBond() {
	if h2o.waiting < 3 {
		return
	}
	select {
	case h2o.hQueue <- 1:
		select {
		case h2o.hQueue <- 1:
			select {
			case h2o.oQueue <- 1:
				h2o.waiting -= 3
			default:
				// release hydrogen atoms and try again
				<-h2o.hQueue
				<-h2o.hQueue
				h2o.waiting -= 2
				h2o.tryBond()
			}
		default:
			// release hydrogen atom and try again
			<-h2o.hQueue
			h2o.waiting--
			h2o.tryBond()
		}
	default:
		// do nothing
	}
}
