package b_Tree

import (
	"errors"
)

var (
	err   error
	queue *Node

	defaultOrder = 3
	minOrder     = 3
	maxOrder     = 20

	order = defaultOrder
)

var (
	errKeyNotFound     = errors.New("key not found")
	errKeyAlreadyExist = errors.New("key already exist")
)

type Record struct {
	value string
}

type Node struct {
	pointers []interface{}
	keys     []int
	parent   *Node
	isLeaf   bool
	numKeys  int
	next     *Node
}

type BPTree struct {
	root *Node
}

func (b *BPTree) insert(key int, value string) error {
	// 确定key在树不存在
	if _, err := b.find(key); err == nil {
		return errKeyAlreadyExist
	} else if err != errKeyNotFound {
		return err
	}

	// 创造记录
	p := makeRecord(value)

	// 若树不存在，创造树
	if b.root == nil {
		b.createNewTree(key, p)
		return nil
	}

	// 找到key最可能所在的叶节点
	leaf := b.findLeaf(key)

	// 若叶节点数量少于规定数量，则插入叶节点
	if leaf.numKeys < order-1 {
		insertIntoLeaf(leaf, key, p)
		return nil
	}

	// 否则进行变换插入叶节点
	b.insertIntoLeafAfterSplitting(leaf, key, p)
}

// insertIntoLeaf 插入叶节点
func insertIntoLeaf(leaf *Node, key int, p *Record) {
	var i, insertPoint int
	// 找到不小于key的插入点
	for insertPoint < leaf.numKeys && leaf.keys[insertPoint] < key {
		insertPoint++
	}

	// 将插入点后面的都后移一位
	for i = leaf.numKeys; i > insertPoint; i-- {
		leaf.keys[i] = leaf.keys[i-1]
		leaf.pointers[i] = leaf.pointers[i-1]
	}

	// 插入记录、key并更新key数量
	leaf.keys[insertPoint] = key
	leaf.pointers[insertPoint] = p
	leaf.numKeys++
	return
}

func makeRecord(value string) *Record {
	return &Record{value: value}
}

// find 根据key找到key的记录
func (b *BPTree) find(key int) (*Record, error) {
	r := b.findLeaf(key)
	if r == nil {
		return nil, errKeyNotFound
	}

	i := 0
	for i = 0; i < r.numKeys; i++ {
		if r.keys[i] == key {
			break
		}
	}

	if i == r.numKeys {
		return nil, errKeyNotFound
	}

	return r.pointers[i].(*Record), nil
}

// findLeaf 找到key最可能所在的叶子节点
func (b *BPTree) findLeaf(key int) *Node {
	r := b.root
	if r == nil {
		return r
	}

	i := 0
	for !r.isLeaf {
		i = 0
		// 找到恰好大于等于key的index
		for i < r.numKeys {
			if key >= r.keys[i] {
				i += 1
			} else {
				break
			}
		}
		// 将递归节点更新为index，知道找到叶节点
		r = r.pointers[i].(*Node)
	}

	return r
}

// createNewTree 以第一个记录创建新树
func (b *BPTree) createNewTree(key int, p *Record) {
	b.root = makeLeaf()
	b.root.keys[0] = key
	b.root.pointers[0] = p
	b.root.pointers[order-1] = nil
	b.root.parent = nil
	b.root.numKeys++
}

func (b *BPTree) insertIntoLeafAfterSplitting(leaf *Node, key int, p *Record) {
	// 找到不小于key的插入点
	var insertIndex int
	for insertIndex < order-1 && leaf.keys[insertIndex] < key {
		insertIndex++
	}

	// 将原叶节点key、pointers复制到临时节点，并给插入点预留位置
	var i, j int
	keys := make([]int, order)
	ps := make([]interface{}, order)
	for i = 0; i < leaf.numKeys; i++ {
		if j == insertIndex {
			j++
		}
		keys[j] = leaf.keys[i]
		ps[j] = leaf.pointers[i]
		j++
	}

	keys[insertIndex] = key
	ps[insertIndex] = p
	leaf.numKeys = 0

	// 将
	split := cutHalf(order - 1)
	for i = 0; i < split; i++ {
		leaf.pointers[i] = ps[i]
		leaf.keys[i] = keys[i]
		leaf.numKeys++
	}

	j = 0
	newLeaf := makeLeaf()
	for i = split; i < order; i++ {
		newLeaf.pointers[i] = ps[i]
		newLeaf.keys[i] = keys[i]
		newLeaf.numKeys++
		j++
	}

	newLeaf.pointers[order-1] = leaf.pointers[order-1]
	leaf.pointers[order-1] = newLeaf

	for i = leaf.numKeys; i < order-1; i++ {
		leaf.pointers[i] = nil
	}
	for i = newLeaf.numKeys; i < order-1; i++ {
		newLeaf.pointers[i] = nil
	}

	var newKey int
	newLeaf.parent = leaf.parent
	newKey = newLeaf.keys[0]

	b.insertIntoParent(leaf, newKey, newLeaf)
}

// insertIntoParent 插入到父节点
func (b *BPTree) insertIntoParent(left *Node, key int, right *Node) {
	parent := left.parent
	// 在只有叶节点情况下，创造跟节点插入
	if parent == nil {
		b.insertInfoNewRoot(left, key, right)
		return
	}
	leftIndex := getLeftIndex(parent, left)

	// 若父节点少于限制的节点数量，直接插入
	if parent.numKeys < order-1 {
		insertIntoNode(parent, leftIndex, key, right)
		return
	}

	b.insertIntoNodeAfterSplitting(parent, leftIndex, key, right)
}

func insertIntoNode(parent *Node, leftIndex int, key int, right *Node) {
	for i := parent.numKeys; i < leftIndex; i-- {
		parent.pointers[i+1] = parent.pointers[i]
		parent.keys[i] = parent.keys[i-1]
	}

	parent.pointers[leftIndex+1] = right
	parent.keys[leftIndex] = key
	parent.numKeys++
}

func getLeftIndex(parent *Node, left *Node) int {
	leftIndex := 0
	for leftIndex <= parent.numKeys && parent.pointers[leftIndex] != left {
		leftIndex += 1
	}
	return leftIndex
}

// insertInfoNewRoot 创建跟节点，并使左右节点置于跟节点下的叶节点
func (b *BPTree) insertInfoNewRoot(left *Node, key int, right *Node) {
	b.root = makeNode()
	b.root.keys[0] = key
	b.root.pointers[0] = left
	b.root.pointers[1] = right
	b.root.numKeys++
	b.root.parent = nil
	left.parent = b.root
	right.parent = b.root
}

func (b *BPTree) insertIntoNodeAfterSplitting(old *Node, leftIndex int, key int, right *Node) {
	tmpPointers := make([]interface{}, order+1)
	tmpKeys := make([]int, order)

	var j int
	for i := 0; i < old.numKeys+1; i++ {
		if j == leftIndex+1 {
			j++
		}
		tmpPointers[j] = old.pointers[i]
	}
}

func cutHalf(length int) int {
	if length%2 == 0 {
		return length / 2
	}
	return length/2 + 1
}

// makeLeaf 创建B+树的叶子节点
func makeLeaf() *Node {
	leaf := makeNode()
	leaf.isLeaf = true
	return leaf
}

// makeNode 创建节点
func makeNode() *Node {
	return &Node{
		pointers: make([]interface{}, order),
		keys:     make([]int, order-1),
		parent:   nil,
		isLeaf:   false,
		numKeys:  0,
		next:     nil,
	}
}

func NewBPTree() *BPTree {
	return &BPTree{}
}
