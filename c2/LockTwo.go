package c2

import "runtime"

type LockTwo struct {
	victim int
}

func (l *LockTwo) Lock(i int) {
	l.victim = i
	for l.victim == i {
		runtime.Gosched()
	}
}

func (l *LockTwo) Unlock(i int) {
}
