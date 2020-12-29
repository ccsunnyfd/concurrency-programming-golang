package recursivemutexv2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// TokenRecursiveMutex Token方式的递归锁 包装一个Mutex,实现可重入
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32 // 这个goroutine 重入的次数
}

// Lock is
func (m *TokenRecursiveMutex) Lock(token int64) {
	// 如果传入的token和持有锁的token一致，说明是递归调用
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 记录下它的token,调用次数加1
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

// Unlock is
func (m *TokenRecursiveMutex) Unlock(token int64) {
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.token, 0)
	m.Mutex.Unlock()
}
