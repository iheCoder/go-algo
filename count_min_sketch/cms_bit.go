package count_min_sketch

import (
	"encoding/binary"
	"errors"
	"hash"
	"hash/fnv"
	"math"
)

// 在CountMinSketch基础上采用bit记录访问次数，以达到最大程度节约内存
// 默认以4bit记录次数，当到达上限后，便不再增加
// 有一个很致命的问题，cms记录无法删除啊

type CMSBitVersion struct {
	// num of hashing func
	d uint
	// size of every hash table
	w uint
	// a matrix, is used to store the cmsRecords
	cmsRecords []*ConcurrentCMSRecords
	hasher     hash.Hash64
}

func newCMSBit(d, w uint) (*CMSBitVersion, error) {
	if d <= 0 || w <= 0 {
		return nil, errors.New("cmsRecords min sketch: values of d and w should both be greater than 0")
	}
	count := make([]*ConcurrentCMSRecords, d)
	for r := uint(0); r < d; r++ {
		count[r] = NewCMSCounter(w)
	}

	s := &CMSBitVersion{
		d:          d,
		w:          w,
		cmsRecords: count,
		hasher:     fnv.New64(),
	}

	return s, nil
}

func NewCMSBitVersion(epsilon, delta float64) (*CMSBitVersion, error) {
	if epsilon <= 0 || epsilon >= 1 {
		return nil, errors.New("countminsketch: value of epsilon should be in range of (0, 1)")
	}
	if delta <= 0 || delta >= 1 {
		return nil, errors.New("countminsketch: value of delta should be in range of (0, 1)")
	}

	w := uint(math.Ceil(math.E / epsilon))
	d := uint(math.Ceil(math.Log(1 / delta)))

	return newCMSBit(d, w)
}

func (s *CMSBitVersion) Update(key string, count uint) {
	s.update([]byte(key), count)
}

func (s *CMSBitVersion) update(key []byte, count uint) {
	for r, c := range s.locations(key) {
		s.cmsRecords[r].Increase(int(c), int(count))
	}
}

func (s *CMSBitVersion) locations(key []byte) []uint {
	locs := make([]uint, s.d)
	a, b := s.bashHashes(key)
	ua := uint(a)
	ub := uint(b)
	for r := uint(0); r < s.d; r++ {
		locs[r] = (ua + ub*r) % s.w
	}

	return locs
}

func (s *CMSBitVersion) bashHashes(key []byte) (uint32, uint32) {
	s.hasher.Reset()
	s.hasher.Write(key)
	sum := s.hasher.Sum(nil)
	upper := sum[:4]
	lower := sum[4:8]
	a := binary.BigEndian.Uint32(lower)
	b := binary.BigEndian.Uint32(upper)
	return a, b
}

func (s *CMSBitVersion) Estimate(key string) int {
	return s.estimate([]byte(key))
}

func (s *CMSBitVersion) estimate(key []byte) int {
	var min int
	var x int
	// 缓存行，多线程
	for r, c := range s.locations(key) {
		x = s.cmsRecords[r].GetCount(int(c))
		if r == 0 || x < min {
			min = x
		}
	}
	return min
}

func (s *CMSBitVersion) Delay2() {
	for _, record := range s.cmsRecords {
		record.Delay2()
	}
}
