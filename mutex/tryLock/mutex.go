package trylock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving    // 从state字段中分出一个饥饿标记
	mutexWaiterShift = iota
)

// Mutex is a mutex with TryLock method
type Mutex struct {
	sync.Mutex
}

// TryLock is
func (m *Mutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
		return false
	}
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}
