package timewheel

// edit from https://github.com/ouqiang/timewheel/blob/master/timewheel.go
import (
	"container/list"
	"time"
)

type job func(s string)

// 计算position、circle有问题。比如position，当为1的时候，pos居然也是1，可不是应该为0嘛？

type task struct {
	delay time.Duration
	// 时间轮需要转动几圈
	circle int
	// task 唯一标识，即使task内容一致也不应该相同
	key  string
	data string
}

// 一个时间轮，轮中每个slot放上任务
type timeWheel struct {
	interval          time.Duration
	ticker            *time.Ticker
	slots             []*list.List
	currentPos        int
	slotNum           int
	doJob             job
	addTaskChannel    chan *task
	removeTaskChannel chan string
	stopChannel       chan bool
	exists            map[string]struct{}
	deleted           map[string]struct{}
}

// initSlots 初始化slots
func (w *timeWheel) initSlots() {
	for i := 0; i < w.slotNum; i++ {
		w.slots[i] = list.New()
	}
}

func NewTimeWheel(interval time.Duration, slotNum int, job job) *timeWheel {
	if interval <= 0 || slotNum <= 0 || job == nil {
		return nil
	}
	tw := &timeWheel{
		interval:          interval,
		ticker:            nil,
		slots:             make([]*list.List, slotNum),
		currentPos:        0,
		slotNum:           slotNum,
		doJob:             job,
		addTaskChannel:    make(chan *task),
		removeTaskChannel: make(chan string),
		stopChannel:       make(chan bool),
		exists:            make(map[string]struct{}),
		deleted:           make(map[string]struct{}),
	}

	tw.initSlots()
	return tw
}

func (w *timeWheel) Start() {
	w.ticker = time.NewTicker(w.interval)
	go w.start()
}

func (w *timeWheel) scanAndRunTask(l *list.List) {
	var next *list.Element
	for e := l.Front(); e != nil; {
		t := e.Value.(*task)
		_, ok := w.deleted[t.key]
		if !ok && t.circle > 0 {
			t.circle--
			e = e.Next()
			continue
		}

		if ok {
			delete(w.deleted, t.key)
		} else {
			go w.doJob(t.data)
			delete(w.exists, t.key)
		}

		next = e.Next()
		l.Remove(e)
		e = next
	}
}

func (w *timeWheel) start() {
	for {
		select {
		case <-w.ticker.C:
			w.tickHandler()
		case t := <-w.addTaskChannel:
			w.addTask(t)
		case key := <-w.removeTaskChannel:
			w.removeTask(key)
		case <-w.stopChannel:
			w.ticker.Stop()
			return
		}
	}
}

func (w *timeWheel) tickHandler() {
	// 获取当前时刻的slot
	l := w.slots[w.currentPos]
	// 扫描该slot中任务是否到了该执行的时刻
	w.scanAndRunTask(l)
	// 更新当前指向位置
	if w.currentPos == w.slotNum-1 {
		w.currentPos = 0
	} else {
		w.currentPos++
	}
}

func (w *timeWheel) addTask(task *task) {
	if len(task.key) == 0 {
		panic("invalid key")
	}

	if task.delay < w.interval {
		panic("delay must larger than internal")
	}

	pos := w.getPos(task.delay)
	circle := w.getCircle(task.delay)
	task.circle = circle

	w.slots[pos].PushBack(task)

	w.exists[task.key] = struct{}{}
}

func (w *timeWheel) removeTask(key string) {
	if len(key) == 0 {
		panic("empty key")
	}

	_, ok := w.exists[key]
	if !ok {
		panic("no such key")
	}

	w.deleted[key] = struct{}{}

	// delete task in exists
	delete(w.exists, key)
}

func (w *timeWheel) getCircle(delay time.Duration) int {
	delaySeconds := int(delay.Seconds())
	internalSeconds := int(w.interval.Seconds())
	circle := (delaySeconds/internalSeconds - 1) / w.slotNum
	return circle
}

func (w *timeWheel) getPos(delay time.Duration) int {
	delaySeconds := int(delay.Seconds())
	internalSeconds := int(w.interval.Seconds())
	// 为什么要加上currentPos呢？
	pos := ((delaySeconds / internalSeconds) - 1) % w.slotNum
	return pos
}

func (w *timeWheel) Stop() {
	w.stopChannel <- true
}

func (w *timeWheel) AddTask(delay time.Duration, key, data string) {
	if delay <= 0 {
		panic("invalid delay")
	}
	w.addTaskChannel <- &task{
		delay: delay,
		key:   key,
		data:  data,
	}
}

func (w *timeWheel) RemoveTask(key string) {
	if len(key) == 0 {
		panic("invalid key")
	}
	w.removeTaskChannel <- key
}
