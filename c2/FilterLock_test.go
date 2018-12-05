package c2

import (
	"sync"
	"testing"
	"time"
)

const N = 10000
const NumThreads = 100

func TestFilterLock(t *testing.T) {

	l := NewFilterLock(NumThreads)
	counter := int64(0)
	var wait sync.WaitGroup
	wait.Add(NumThreads)
	for t := int32(0); t < NumThreads; t++ {
		go func(t int32) {
			defer wait.Done()
			for i := 0; i < N; i++ {
				l.Lock(t)
				counter = counter + 1
				l.Unlock(t)
			}
		}(t)
	}

	wait.Wait()
	time.Sleep(time.Second)
	if counter != N*NumThreads {
		t.Errorf("counter is %d, should be %d", counter, N*NumThreads)
	}
}

func TestFilterLockVSSyncMutex(t *testing.T) {

	var l sync.Mutex
	counter := int64(0)
	var wait sync.WaitGroup
	wait.Add(NumThreads)
	for t := int32(0); t < NumThreads; t++ {
		go func(t int32) {
			defer wait.Done()
			for i := 0; i < N; i++ {
				l.Lock()
				counter = counter + 1
				l.Unlock()
			}
		}(t)
	}

	wait.Wait()
	time.Sleep(time.Second)
	if counter != N*NumThreads {
		t.Errorf("counter is %d, should be %d", counter, N*NumThreads)
	}
}
