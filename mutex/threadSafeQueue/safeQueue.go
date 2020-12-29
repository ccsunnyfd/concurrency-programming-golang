package threadsafequeue

import "sync"

// SafeQueue is
type SafeQueue struct {
	mu   sync.Mutex
	data []interface{}
}

// NewSafeQueue is
func NewSafeQueue(n int) (q *SafeQueue) {
	return &SafeQueue{data: make([]interface{}, 0, n)}
}

//Enqueue is
func (q *SafeQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
}

// Dequeue is
func (q *SafeQueue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) <= 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
