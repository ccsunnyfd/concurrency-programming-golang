package hackinnerstatistics

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

// Mutex is
type Mutex struct {
	sync.Mutex
}

// Count is
func (m *Mutex) Count() int {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return int(state&mutexLocked + state>>mutexWaiterShift)
}

// IsWoken is
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// IsLocked is
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// IsStarving is
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
