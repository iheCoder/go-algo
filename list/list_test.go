package list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestListUse(t *testing.T) {
	l := list.New()
	back := l.Back()
	fmt.Println(back.Value)
	l.PushBack("first")
	e1 := l.PushBack(100)
	e := l.PushFront("ele")
	l.InsertBefore("front", e)
	l.InsertAfter("back", e)
	l.Remove(e)
	l.MoveToFront(e1)
}

func BenchmarkWillConflictIfPushBackAndFront(b *testing.B) {
	l := list.New()
	for i := 0; i < b.N; i++ {
		//go pushFront(l)
		go pushBack(l)
		go moveFront(l)
		//go remove(l)
	}
}

//func remove(l *list.List) {
//	e:=l.Front()
//	l.Remove(e)
//}

var value = 0

func pushBack(l *list.List) {
	l.PushBack(value)
	value++
}

func pushFront(l *list.List) {
	l.PushFront(value)
	value++
}

func moveFront(l *list.List) {
	back := l.Back()
	l.MoveToFront(back)
}
