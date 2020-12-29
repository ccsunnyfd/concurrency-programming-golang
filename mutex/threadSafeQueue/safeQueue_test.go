package threadsafequeue

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSafeQueue(t *testing.T) {
	var queue = NewSafeQueue(100)
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		i := i
		go func(int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			if i%2 == 0 {
				queue.Enqueue(i)
			} else {
				queue.Enqueue(fmt.Sprintf("whatever%d", i))
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	t.Log(queue)
	t.Logf("1st element: %v", queue.Dequeue())
	t.Logf("2nd element: %v", queue.Dequeue())
}
