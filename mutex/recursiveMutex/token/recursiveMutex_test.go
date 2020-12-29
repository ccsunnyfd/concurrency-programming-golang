package recursivemutexv2

import (
	"fmt"
	"testing"
)

func foo(l *TokenRecursiveMutex) {
	fmt.Println("in foo")
	l.Lock(1234)
	bar(l)
	l.Unlock(1234)
}

func bar(l *TokenRecursiveMutex) {
	l.Lock(1234)
	// l.Lock(5678)
	fmt.Println("in bar")
	// l.Unlock(5678)
	l.Unlock(1234)
}

func TestRecursiveMutexV2(t *testing.T) {
	l := &TokenRecursiveMutex{}
	foo(l)
}
