package count_min_sketch

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func TestBashHashes(t *testing.T) {
	x := "hello"
	s, err := NewCMSWithEstimates(0.1, 0.8)
	if err != nil {
		t.Fatal(err)
	}
	s.Update(x, 1)
	s.Estimate("hello")
}

func TestCMSBitUse(t *testing.T) {
	cms, err := NewCMSBitVersion(0.1, 0.8)
	if err != nil {
		t.Fatal(err)
	}
	cms.Update("a", 1)
	if cms.Estimate("a") != 1 {
		t.Fatal("a not cmsRecords 1")
	}

	cms.Update("a", 10)
	if cms.Estimate("a") != 11 {
		t.Fatal("a not cmsRecords 11")
	}
}

func TestConcurrentPerformanceOfCMSBit(t *testing.T) {
	cms, err := NewCMSBitVersion(0.1, 0.8)
	if err != nil {
		t.Fatal(err)
	}
	count := 100000
	wg := &sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		x := rand.Intn(100)
		if x%3 == 0 {
			go updateWg(cms, strconv.Itoa(x), wg)
		} else {
			go estimateWg(cms, strconv.Itoa(x), wg)
		}
	}
	wg.Wait()
}

func updateWg(cms *CMSBitVersion, key string, wg *sync.WaitGroup) {
	cms.Update(key, 1)
	wg.Done()
}

func estimateWg(cms *CMSBitVersion, key string, wg *sync.WaitGroup) {
	cms.Estimate(key)
	wg.Done()
}
