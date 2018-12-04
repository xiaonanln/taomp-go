package c2

import "runtime"

type PetersonLock struct {
	flag   [2]bool
	victim int
}

func (l *PetersonLock) Lock(i int) {
	j := 1 - i
	l.flag[i] = true
	l.victim = i
	for l.flag[j] && l.victim == i {
		runtime.Gosched()
	}
}

func (l *PetersonLock) Unlock(i int) {
	l.flag[i] = false
}
