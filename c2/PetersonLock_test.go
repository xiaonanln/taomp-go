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
	var c1, c2 bool
	wait.Add(2)
	go func() {
		defer wait.Done()
		for i := 0; i < N; i++ {
			l.Lock(0)
			if c2 {
				println("bad")
			}
			c1 = true
			if c2 {
				println("bad")
			}
			counter = counter + 1
			c1 = false
			l.Unlock(0)
		}
	}()

	go func() {
		defer wait.Done()
		for i := 0; i < N; i++ {
			l.Lock(1)
			if c1 {
				println("bad")
			}
			c2 = true
			if c1 {
				println("bad")
			}
			counter = counter + 1
			c2 = false
			l.Unlock(1)
		}
	}()

	wait.Wait()
	time.Sleep(time.Second)
	if counter != N*2 {
		t.Errorf("counter is %d, should be %d", counter, N*2)
	}

	println(c1, c2)
}
