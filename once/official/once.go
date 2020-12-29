package myonce

import (
	"sync"
	"sync/atomic"
)

// Once is
type Once struct {
	sync.Mutex
	done uint32
}

// Do is
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

// slowPath is
func (o *Once) doSlow(f func()) {
	o.Lock()
	defer o.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
