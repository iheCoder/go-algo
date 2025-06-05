package heavy_keeper

import (
	"encoding/binary"
	"github.com/OneOfOne/xxhash"
	"math"
	"math/rand"
	"sync"
)

type bucket struct {
	fingerprint uint64
	count       uint64
	mu          sync.Mutex
}

type TopK struct {
	k     uint32
	width uint32
	depth uint32
	decay float64
	items chan []byte

	seed    int
	buckets [][]bucket
	minHeap *Heap
	wg      *sync.WaitGroup
}

func (b *bucket) Get() (uint64, uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.fingerprint, b.count
}

func (b *bucket) Set(fingerprint uint64, count uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.fingerprint = fingerprint
	b.count = count
}

func (b *bucket) Inc(val uint64) uint64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.count += val
	return b.count
}

func (tk *TopK) Wait() {
	close(tk.items)
	tk.wg.Wait()
}

func (tk *TopK) AddBytes(item []byte) {
	tk.items <- item
}

func (tk *TopK) CountBytes(item []byte) (uint64, bool) {
	if id, exists := tk.minHeap.Find(item); exists {
		return tk.minHeap.Nodes[id].Count, true
	}

	return 0, false
}

func (tk *TopK) QueryBytes(item []byte) bool {
	_, exists := tk.minHeap.Find(item)
	return exists
}

func (tk *TopK) List() Nodes {
	return tk.minHeap.Sorted()
}

func (tk *TopK) jobAdder() {
	for item := range tk.items {
		itemFp := xxhash.Checksum64S(item, uint64(tk.seed))

		var maxCount uint64
		itemHeapIdx, exists := tk.minHeap.Find(item)
		heapMin := tk.minHeap.Min()

		for i := uint32(0); i < tk.depth; i++ {
			bi := make([]byte, 4)
			binary.BigEndian.PutUint32(bi, i)

			// get the bucket number
			bucketNum := xxhash.Checksum64S(append(item, bi...), uint64(tk.seed)) % uint64(tk.width)

			tk.buckets[i][bucketNum].mu.Lock()

			// get the fingerprint and count of the bucket
			fpBucket, count := tk.buckets[i][bucketNum].fingerprint, tk.buckets[i][bucketNum].count

			// if the bucket is empty, set the fingerprint and count
			// if the fingerprint is the same as the item's fingerprint, increment the count
			// if the fingerprint is different, decrement the count
			if count == 0 {
				tk.buckets[i][bucketNum].fingerprint = itemFp
				tk.buckets[i][bucketNum].count = 1
				maxCount = maxUint64(maxCount, 1)
			} else if itemFp == fpBucket {
				if exists || count <= heapMin {
					tk.buckets[i][bucketNum].count++
					maxCount = maxUint64(maxCount, tk.buckets[i][bucketNum].count)
				}
			} else {
				decay := math.Pow(tk.decay, float64(count))
				if rand.Float64() < decay {
					tk.buckets[i][bucketNum].count--
					if tk.buckets[i][bucketNum].count == 0 {
						tk.buckets[i][bucketNum].fingerprint = itemFp
						tk.buckets[i][bucketNum].count = 1
						maxCount = maxUint64(maxCount, 1)
					}
				}
			}

			tk.buckets[i][bucketNum].mu.Unlock()
		}

		if exists {
			tk.minHeap.Fix(itemHeapIdx, maxCount)
		} else {
			tk.minHeap.Add(Node{Item: item, Count: maxCount})
		}
	}
}

func maxUint64(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}
