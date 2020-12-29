package rwmap

import (
	"sync"
	"testing"
)

func ltThree(k, v int) bool {
	return k+v < 7
}

func TestRWMap(t *testing.T) {
	var wg sync.WaitGroup

	var m = NewRWMap(10) // 初始化一个map
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			m.Set(i, i) //设置key
		}
		m.Each(ltThree)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			_, ok := m.Get(2) //访问这个map
			if !ok {
				t.Log("key 2's corresponding value is not ready yet!")
			}
		}
	}()
	wg.Wait()
}
