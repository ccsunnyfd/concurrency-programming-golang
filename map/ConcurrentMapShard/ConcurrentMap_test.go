package concurrentmapshard

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func ltThree(k, v int) bool {
	return k+v < 7
}

func TestConcurrentMapShard(t *testing.T) {
	var wg sync.WaitGroup

	var m = New() // 初始化一个map
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Microsecond)
			m.Set(strconv.Itoa(i), i) //设置key
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Microsecond)
			v, ok := m.Get(strconv.Itoa(i)) //访问这个map
			if !ok {
				t.Logf("key %d's corresponding value is not ready yet!", i)
			} else {
				t.Logf("key %d's corresponding value is %d!", i, v)
			}
		}
	}()
	wg.Wait()
}
