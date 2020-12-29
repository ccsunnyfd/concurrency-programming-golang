package myonce2

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 值是3.0或者0.0的一个数据结构
var threeOnce struct {
	Once
	v float32
}

func testInitError() error {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rd.Intn(10)
	if r > 4 {
		fmt.Printf("threeOnce initing, rand num: %d\n", r)
		threeOnce.v = 3.0
		return nil
	}
	return fmt.Errorf("init error, rand num: %d", r)
}

// 返回此数据结构的值，如果还没有初始化为3.0，则初始化
func three() (float32, error) {
	if err := threeOnce.Do(testInitError); err != nil {
		return 0.0, err
	}
	return threeOnce.v, nil
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		go func() {
			defer wg.Done()
			if v, err := three(); err != nil {
				t.Log(err)
			} else {
				t.Log(v)
			}
		}()
	}
	wg.Wait()
}
