package kvstore

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// KVStoreService is
type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.RWMutex
}

// NewKVStoreService is
func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

// Get is
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

// Set is
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

// Watch is
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}

	//return nil
}
