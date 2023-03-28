package count_min_sketch

import (
	"encoding/binary"
	"errors"
	"hash"
	"hash/fnv"
	"math"
)

type CountMinSketch struct {
	// num of hashing func
	d uint
	// size of every hash table
	w uint
	// a matrix, is used to store the count
	count  [][]uint64
	hasher hash.Hash64
}

func newCMS(d, w uint) (s *CountMinSketch, err error) {
	if d <= 0 || w <= 0 {
		return nil, errors.New("cmsRecords min sketch: values of d and w should both be greater than 0")
	}

	s = &CountMinSketch{
		d:      d,
		w:      w,
		hasher: fnv.New64(),
	}
	s.count = make([][]uint64, d)
	for r := uint(0); r < d; r++ {
		s.count[r] = make([]uint64, w)
	}

	return s, nil
}

// NewCMSWithEstimates 在1-delta概率下总误差小于epsilon
func NewCMSWithEstimates(epsilon, delta float64) (*CountMinSketch, error) {
	if epsilon <= 0 || epsilon >= 1 {
		return nil, errors.New("countminsketch: value of epsilon should be in range of (0, 1)")
	}
	if delta <= 0 || delta >= 1 {
		return nil, errors.New("countminsketch: value of delta should be in range of (0, 1)")
	}

	w := uint(math.Ceil(math.E / epsilon))
	d := uint(math.Ceil(math.Log(1 / delta)))

	return newCMS(d, w)
}

func (s *CountMinSketch) Update(key string, count uint64) {
	s.update([]byte(key), count)
}

func (s *CountMinSketch) update(key []byte, count uint64) {
	for r, c := range s.locations(key) {
		s.count[r][c] += count
	}
}

func (s *CountMinSketch) locations(key []byte) []uint {
	locs := make([]uint, s.d)
	a, b := s.bashHashes(key)
	ua := uint(a)
	ub := uint(b)
	for r := uint(0); r < s.d; r++ {
		locs[r] = (ua + ub*r) % s.w
	}

	return locs
}

func (s *CountMinSketch) bashHashes(key []byte) (uint32, uint32) {
	s.hasher.Reset()
	s.hasher.Write(key)
	sum := s.hasher.Sum(nil)
	upper := sum[:4]
	lower := sum[4:8]
	a := binary.BigEndian.Uint32(lower)
	b := binary.BigEndian.Uint32(upper)
	return a, b
}

func (s *CountMinSketch) Estimate(key string) uint64 {
	return s.estimate([]byte(key))
}

func (s *CountMinSketch) estimate(key []byte) uint64 {
	var min uint64
	for r, c := range s.locations(key) {
		if r == 0 || s.count[r][c] < min {
			min = s.count[r][c]
		}
	}
	return min
}
