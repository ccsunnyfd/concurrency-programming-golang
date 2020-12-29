package recursivemutexv1

import (
	"fmt"
	"sync"
	"testing"
)

func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func TestRecursiveMutexV1(t *testing.T) {
	l := &RecursiveMutex{}
	foo(l)
}
