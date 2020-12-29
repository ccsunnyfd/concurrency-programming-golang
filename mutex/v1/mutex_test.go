package v1

import (
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var counter = newCounter()

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				time.Sleep(time.Millisecond * 1 / 2)
				counter.Incr(1)
			}
		}()
	}
	wg.Wait()
	counter.Close()
	t.Logf("Total count: %d", counter.Count())
}
