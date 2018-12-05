package c2

import (
	"runtime"
	"sync/atomic"
)

type FilterLock struct {
	threadNum int32
	level     []int32
	victim    []int32
}

func NewFilterLock(threadNum int) *FilterLock {
	l := &FilterLock{
		threadNum: int32(threadNum),
		level:     make([]int32, threadNum),
		victim:    make([]int32, threadNum),
	}
	return l
}

func (l *FilterLock) Lock(me int32) {
	for i := int32(1); i < l.threadNum; i++ {
		atomic.StoreInt32(&l.level[me], i)
		atomic.StoreInt32(&l.victim[i], me)
		for {
			blocked := false
			for k := int32(0); k < l.threadNum; k++ {
				if k != me && atomic.LoadInt32(&l.level[k]) >= i && atomic.LoadInt32(&l.victim[i]) == me {
					blocked = true
					break
				}
			}
			if !blocked {
				break
			}
			runtime.Gosched()
		}
	}
}

func (l *FilterLock) Unlock(me int32) {
	atomic.StoreInt32(&l.level[me], 0)
}
