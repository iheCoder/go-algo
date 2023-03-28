package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

// mutex 不行，难以保证同一个方法不会连续抢到
// chan int 1也不行。要有一个获取锁解锁的过程，可是chan本身无法做到
// two mutex。thread1 lock m1, unlock m2 finally, thread2 lock m2, unlock m1 finally. 这不行啊
// two chan是可以的，当然mutex也能做到，一个mutex也能做到

func TestPrintByTurnChBlock(t *testing.T) {
	for i := 0; i < 13; i++ {
		go printNumberCh()
		<-numCh
		go printAlphabetCh()
		<-alpCh
	}
}

var number int
var alphabet = 'a'
var numCh = make(chan int)
var alpCh = make(chan int)

func printNumberCh() {
	printNumber()
	numCh <- 1
}

func printNumber() {
	number++
	fmt.Print(number)
	number++
	fmt.Print(number)
}

func printAlphabetCh() {
	printAlphabet()
	alpCh <- 1
}

func printAlphabet() {
	fmt.Printf("%c", alphabet)
	alphabet++
	fmt.Printf("%c", alphabet)
	alphabet++
}

var oneMu = &sync.Mutex{}

func TestPrintByTurnOneMutex(t *testing.T) {
	for i := 0; i < 13; i++ {
		oneMu.Lock()
		go printNumberOneMu()
		oneMu.Lock()
		go printAlphabetOneMu()
	}
}

func printNumberOneMu() {
	printNumber()
	oneMu.Unlock()
}

func printAlphabetOneMu() {
	printAlphabet()
	oneMu.Unlock()
}

func TestPrintByTurnInnerCh(t *testing.T) {
	letter, num := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}

				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				num <- true
			}

		}
	}(&wait)
	num <- true
	wait.Wait()
}
