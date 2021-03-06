package official

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var counter counter
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				// time.Sleep(time.Millisecond * 1)
				counter.Incr(1)
			}
		}()
	}
	wg.Wait()
	t.Logf("Total count: %d", counter.Count())
}
