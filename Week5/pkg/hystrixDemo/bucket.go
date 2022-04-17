package hystrixDemo

import (
	"sync"
	"time"
)

type Bucket struct {
	sync.RWMutex
	Total     int
	Fail      int
	TimeStamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		TimeStamp: time.Now(),
	}
}

func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()
	if !result {
		b.Fail++
	}
	b.Total++
}
