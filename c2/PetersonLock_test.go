package c2

import (
	"sync"
	"testing"
	"time"
)

func TestPetersonLock(t *testing.T) {
	var l PetersonLock
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
	time.Sleep(time.Second)
	if counter != N*2 {
		t.Errorf("counter is %d, should be %d", counter, N*2)
	}
}

func TestMutex(t *testing.T) {
	var l sync.Mutex
	counter := int64(0)
	var wait sync.WaitGroup
	const N = 1000000
	wait.Add(2)
	go func() {
		defer wait.Done()
		for i := 0; i < N; i++ {
			l.Lock()
			counter = counter + 1
			l.Unlock()
		}
	}()

	go func() {
		defer wait.Done()
		for i := 0; i < N; i++ {
			l.Lock()
			counter = counter + 1
			l.Unlock()
		}
	}()

	wait.Wait()
	time.Sleep(time.Second)
	if counter != N*2 {
		t.Errorf("counter is %d, should be %d", counter, N*2)
	}

}
