package ReadWriteMutex

import (
	"sync"
	"time"
)

type ReadWriteMutex struct {
	readersCount int
	readersLock  sync.Mutex
	globalLock   sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCount++
	if rw.readersCount == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCount--
	if rw.readersCount == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}
