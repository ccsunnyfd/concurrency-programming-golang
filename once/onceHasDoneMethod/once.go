package oncehasdonemethod

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Once is
type Once struct {
	sync.Once
}

// Done is
func (o *Once) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}
