package skipTable

import "math/rand"

// copy from https://github.com/wangzheng0822/algo/blob/master/java/17_skiplist/SkipList.java

const (
	skipListP  = 0.5
	maxLevelWz = 16
)

type skipListWz struct {
	levelCount int64
	head       *NodeWz
}

type NodeWz struct {
	data     int64
	forwards []*NodeWz
	level    int64
}

func (l *skipListWz) find(val int64) *NodeWz {
	p := l.head
	for i := l.levelCount - 1; i >= 0; i-- {
		for p.forwards[i] != nil && p.forwards[i].data < val {
			p = p.forwards[i]
		}
	}
}

func NewNodeWz() *NodeWz {
	return &NodeWz{
		data:     -1,
		forwards: make([]*NodeWz, maxLevelWz),
		level:    0,
	}
}

func randomLevel() int64 {
	level := 1
	for rand.Float64() < skipListP && level < maxLevel {
		level++
	}
	return int64(level)
}

func NewSkipListWz() *skipListWz {
	return &skipListWz{
		levelCount: 1,
	}
}
