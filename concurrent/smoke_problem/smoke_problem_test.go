package smoke_problem

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"testing"
	"time"
)

var (
	sem1        = semaphore.NewWeighted(1)
	sem2        = semaphore.NewWeighted(1)
	sem3        = semaphore.NewWeighted(1)
	providerSem = semaphore.NewWeighted(1)
)

func TestSem(t *testing.T) {
	sem1.Acquire(context.Background(), 1)
	sem2.Acquire(context.Background(), 1)
	sem3.Acquire(context.Background(), 1)
	providerSem.Acquire(context.Background(), 1)

	go provideSem()
	go smoker1Sem()
	go smoker2Sem()
	go smoker3Sem()

	time.Sleep(time.Second)
}

func provideSem() {
	var i int
	for {
		switch i {
		case 0:
			sem1.Release(1)
		case 1:
			sem2.Release(1)
		case 2:
			sem3.Release(1)
		}
		i = (i + 1) & 3
		fmt.Println(i)
		providerSem.Acquire(context.Background(), 1)
	}
}

func smoker1Sem() {
	for {
		sem1.Acquire(context.Background(), 1)
		fmt.Println("smoker1 smoking")
		providerSem.Release(1)
	}
}

func smoker2Sem() {
	for {
		sem2.Acquire(context.Background(), 1)
		fmt.Println("smoker2 smoking")
		providerSem.Release(1)
	}
}

func smoker3Sem() {
	for {
		sem3.Acquire(context.Background(), 1)
		fmt.Println("smoker3 smoking")
		providerSem.Release(1)
	}
}
