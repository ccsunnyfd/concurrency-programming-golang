package v2

import "sync/atomic"

// Mutex is a mocked Mutex implementation version 1
type Mutex struct {
	state int32      // if locker is being held
	sema  *semaphore // signal for block
}

func newMutex() Mutex {
	return Mutex{
		sema: newSemaphore(),
	}
}

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexWaiterShift = iota
)

// Lock implements interface Locker's Lock() function
func (m *Mutex) Lock() {
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	awoke := false // not in for loop, so must be first called by Lock(), should be a new coming Goroutine
	for {
		old := m.state
		locked := old&mutexLocked == mutexLocked
		new := old | mutexLocked // if not locked, should set mutexLocked bit
		if locked {
			new = old + 1<<mutexWaiterShift // waiters number increase by 1
		}
		if awoke {
			// if this goroutine is awakened, should clear mutexWoken flag bit
			new &^= mutexWoken
		}

		if atomic.CompareAndSwapInt32(&m.state, old, new) { // if set new state successfully
			if !locked { // if former state is unlocked, means it's a grabLocker success operation
				break
			}
			m.sema.semacquire()
			awoke = true
		}

	}
}

// Unlock implements interface Locker's Unlock() function
func (m *Mutex) Unlock() {

	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked) //去掉锁标志
	if (new+mutexLocked)&mutexLocked == 0 {        //本来就没有加锁
		panic("sync: unlock of unlocked mutex")
	}

	old := new
	for {
		if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 { // 没有等待者，或者有唤醒的waiter，或者锁原来已加锁
			return
		}
		new = (old - 1<<mutexWaiterShift) | mutexWoken // 新状态，准备唤醒goroutine，并设置唤醒标志
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			m.sema.semarelease()
			return
		}
		old = m.state
	}
}

// Close chan
func (m *Mutex) Close() {
	m.sema.close()
}
