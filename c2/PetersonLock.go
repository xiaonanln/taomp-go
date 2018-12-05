package c2

import (
	"runtime"
	"sync/atomic"
)

type PetersonLock struct {
	flag   [2]int32
	victim int32
}

func (l *PetersonLock) Lock(i int32) {
	j := 1 - i
	atomic.StoreInt32(&l.flag[i], 1)
	atomic.StoreInt32(&l.victim, i)
	for atomic.LoadInt32(&l.flag[j]) == 1 && atomic.LoadInt32(&l.victim) == i {
		runtime.Gosched()
	}
}

func (l *PetersonLock) Unlock(i int32) {
	atomic.StoreInt32(&l.flag[i], 0)
}
