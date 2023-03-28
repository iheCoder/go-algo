package timewheel

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
	"time"
)

var now time.Time

func jobPrint(data string) {
	fmt.Printf("data: %s time since start: %s \n", data, time.Since(now).String())
}

// slot 5, put 1s, 5s, 8s
func TestTWUse(t *testing.T) {
	slotNum := 5
	w := NewTimeWheel(time.Second, slotNum, jobPrint)
	w.Start()

	now = time.Now()
	w.AddTask(time.Second, "1", "1s")
	w.AddTask(5*time.Second, "5", "5s")
	w.AddTask(8*time.Second, "8", "8s")
	w.AddTask(6*time.Second, "6", "6s")
	w.RemoveTask("5")

	time.Sleep(15 * time.Second)
}

func TestTWMem(t *testing.T) {
	slotNum := 5
	w := NewTimeWheel(time.Second, slotNum, jobPrint)
	w.Start()
	go printMemOcc()
	for j := 0; j < total; j++ {
		d := rand.Intn(timeLimit) + 1
		addTask(w, time.Duration(d)*time.Second)
	}

	time.Sleep(time.Duration(21) * time.Second)
}

func TestTimer(t *testing.T) {
	go printMemOcc()
	for j := 0; j < total; j++ {
		d := rand.Intn(timeLimit) + 1
		go doTimer(time.Duration(d) * time.Second)
	}
	time.Sleep(21 * time.Second)
}

func doTimer(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
	jobPrint(strconv.Itoa(i))
	i++
}

const total = 1000
const timeLimit = 20

var i int

func addTask(tw *timeWheel, t time.Duration) {
	tw.AddTask(t, strconv.Itoa(i), strconv.Itoa(i))
	i++
}

var ms runtime.MemStats
var moT = time.NewTicker(100 * time.Millisecond)

func printMemOcc() {
	for range moT.C {
		runtime.ReadMemStats(&ms)
		fmt.Println(toMB(ms.Alloc))
	}
}

func toMB(b uint64) uint64 {
	return b / 1024
}
