package c2

import "runtime"

type LockOne struct {
	flag [2]bool
}

func (l *LockOne) Lock(i int) {
	j := 1 - i
	l.flag[i] = true
	for l.flag[j] {
		runtime.Gosched()
	}
}

func (l *LockOne) Unlock(i int) {
	l.flag[i] = false
}
