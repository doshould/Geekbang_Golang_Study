package hystrixDemo

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type RollingWindow struct {
	sync.RWMutex
	broke         bool
	size          int
	buckets       []*Bucket
	reqThreshold  int     //触发熔断的请求总数阈值
	failThreshold float64 //触发熔断的失败率阈值
	lastBrokeTime time.Time
	brokeTimeGap  time.Duration
}

func NewRollingWindow(size int, reqThreshold int, failThreshold float64, brokeTimeGap time.Duration) *RollingWindow {
	return &RollingWindow{
		size:          size,
		buckets:       make([]*Bucket, 0, size),
		reqThreshold:  reqThreshold,
		failThreshold: failThreshold,
		brokeTimeGap:  brokeTimeGap,
	}
}

func (r *RollingWindow) AppendBucket() {
	r.Lock()
	defer r.Unlock()
	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.size+1) {
		r.buckets = r.buckets[1:]
	}
}

func (r *RollingWindow) GetBucket() *Bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

func (r *RollingWindow) RecordReqRequest(result bool) {
	r.GetBucket().Record(result)
}

func (r *RollingWindow) ShowAllBucket() {
	for _, v := range r.buckets {
		fmt.Printf("id: [%v] | total: [%d] | failed: [%d]\\n", v.TimeStamp, v.Total, v.Fail)
	}
}

func (r *RollingWindow) Start() {
	go func() {
		for {
			r.AppendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (r *RollingWindow) BrokeJudge() bool {
	r.Lock()
	defer r.Unlock()
	total := 0
	fail := 0
	for _, v := range r.buckets {
		total += v.Total
		fail += v.Fail
	}
	if float64(fail)/float64(total) > r.failThreshold && total > r.reqThreshold {
		return true
	}
	return false
}

func (r *RollingWindow) OverBrokenTimeGap() bool {
	return time.Since(r.lastBrokeTime) > r.brokeTimeGap
}

func (r *RollingWindow) Monitor() {
	go func() {
		for {
			if r.broke {
				if r.OverBrokenTimeGap() {
					r.Lock()
					r.broke = false
					r.Unlock()
				}
				continue
			}
			if r.BrokeJudge() {
				r.Lock()
				r.broke = true
				r.lastBrokeTime = time.Now()
				r.Unlock()
			}
		}
	}()
}

func (r *RollingWindow) ShowStatus() {
	go func() {
		for {
			log.Println(r.broke)
			time.Sleep(time.Second)
		}
	}()
}

func (r *RollingWindow) Broken() bool {
	return r.broke
}
