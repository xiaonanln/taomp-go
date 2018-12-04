package c2

import (
	"sync"
	"testing"
)

func TestLockTwo(t *testing.T) {
	var l LockTwo
	counter := int64(0)
	var wait sync.WaitGroup
	const N = 1000000
	wait.Add(2)
	go func() {
		defer wait.Done()
		for i := 0; i < N; i++ {
			l.Lock(0)
			counter = counter + 1
			l.Unlock(0)
		}
	}()

	go func() {
		defer wait.Done()
		for i := 0; i < N; i++ {
			l.Lock(1)
			counter = counter + 1
			l.Unlock(1)
		}
	}()

	wait.Wait()
	if counter != N*2 {
		t.Errorf("counter is %d, should be %d", counter, N*2)
	}

}
