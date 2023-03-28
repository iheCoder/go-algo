package skipTable

// ref: 1. https://www.jianshu.com/p/9d8296562806
// 		2. https://blog.csdn.net/qq_42956653/article/details/122017851
const (
	maxLevel = 16
)

type Element struct {
	Member string
	Score  int64
}

type Level struct {
	forward *Nodecp
	span    int64
}

type Nodecp struct {
	Element
	backward *Nodecp
	level    []*Level
}

type skipList struct {
	header *Nodecp
	tail   *Nodecp
	length int64
	level  int16
}

func (l *skipList) insert(e Element) *Nodecp {
	updatePre := make([]*Nodecp, maxLevel)
	rankPre := make([]int64, maxLevel)
	node := l.header

	for i := l.level - 1; i >= 0; i-- {
		if i!=l.level-1 {
			rankPre[i] =rankPre[i+1]
		}

		if node.level[i] != nil {
			for (node.level[i].forward!=nil && ) {
				
			}
		}
	}
}
