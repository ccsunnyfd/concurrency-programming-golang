package v1

import "sync/atomic"

// xadd
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := atomic.LoadInt32(val) // 如果不是原子操作使用--race的时候会发现data race，但是因为自旋也没有关系
		// v := *val  // 官方版本
		if cas(val, v, v+delta) {
			return v + delta
		}
	}
	// panic("unreached")
}

// Mutex is a mocked Mutex implementation version 1
type Mutex struct {
	key  int32      // if locker is being held
	sema *semaphore // signal for block
}

func newMutex() Mutex {
	return Mutex{
		sema: newSemaphore(),
	}
}

// Lock implements interface Locker's Lock() function
func (m *Mutex) Lock() {
	if xadd(&m.key, 1) == 1 {
		return
	}
	m.sema.semacquire()
}

// Unlock implements interface Locker's Unlock() function
func (m *Mutex) Unlock() {
	if xadd(&m.key, -1) == 0 {
		return
	}
	m.sema.semarelease()
}

// Close chan
func (m *Mutex) Close() {
	m.sema.close()
}
