package myonce2

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
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 { // fast path
		return nil
	}
	return o.doSlow(f)
}

// slowPath is
func (o *Once) doSlow(f func() error) error {
	o.Lock()
	defer o.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
