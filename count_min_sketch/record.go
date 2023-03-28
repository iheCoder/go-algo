package count_min_sketch

import "sync"

type cmsRecord struct {
	r    uint64
	lock *sync.RWMutex
}

func (c *cmsRecord) increase(mask, v uint64) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.r&mask != mask {
		c.r += v
		return true
	}
	return false
}

func (c *cmsRecord) getCount(mask uint64, offset int) int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	r := c.r
	return int(r & mask >> offset)
}

func (c *cmsRecord) reset(mask uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.r = c.r &^ mask
}

func (c *cmsRecord) delay2() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.r &= delay2Mask
	c.r >>= 1
}

// 用于count min sketch的bit操作。将一定区间作为频率记录，进行新增、重置、获取
type ConcurrentCMSRecords struct {
	records []*cmsRecord
}

const (
	countMax    = 15
	bitCountMax = 16
	bitUnit     = 4
	delay2Mask  = 0xeeeeeeeeeeeeeee
)

func NewCMSCounter(size uint) *ConcurrentCMSRecords {
	l := (size-1)>>bitUnit + 1
	records := make([]*cmsRecord, l)
	for i := uint(0); i < l; i++ {
		records[i] = &cmsRecord{
			r:    0,
			lock: &sync.RWMutex{},
		}
	}
	r := &ConcurrentCMSRecords{records: records}
	return r
}

func (b *ConcurrentCMSRecords) Increase(index, count int) bool {
	if index-1 > bitCountMax*len(b.records) || index < 0 {
		panic("index invalid")
	}
	if count > countMax || count < 0 {
		panic("cmsRecords invalid")
	}

	i := index >> bitUnit
	offset := (index - (i << bitUnit)) << 2
	mask := uint64(0xf) << offset
	v := uint64(count << offset)

	return b.records[i].increase(mask, v)
}

func (b *ConcurrentCMSRecords) GetCount(index int) int {
	if index-1 > bitCountMax*len(b.records) || index < 0 {
		panic("index invalid")
	}
	i := index >> bitUnit
	offset := (index - (i << bitUnit)) << 2
	mask := uint64(0xf) << offset

	return b.records[i].getCount(mask, offset)
}

func (b *ConcurrentCMSRecords) ResetByIndex(index int) {
	if index-1 > bitCountMax*len(b.records) || index < 0 {
		panic("index invalid")
	}

	i := index >> bitUnit
	offset := (index - (i << bitUnit)) << 2
	mask := uint64(0xf) << offset

	b.records[i].reset(mask)
}

// Delay2 将所有访问次数都衰减除以2
// 最容易想到的是先用mask取到每4位，然后让这4位右移，再mask与，最后将所有结果与
// 第二种方式是先mask一个前3位为1，最后为0的数字。很显然第二种更好
func (b *ConcurrentCMSRecords) Delay2() {
	l := len(b.records)
	for i := 0; i < l; i++ {
		b.records[i].delay2()
	}
}
