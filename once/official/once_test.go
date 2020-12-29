package myonce

import (
	"fmt"
	"testing"
)

// 值是3.0或者0.0的一个数据结构
var threeOnce struct {
	Once
	v float32
}

// 返回此数据结构的值，如果还没有初始化为3.0，则初始化
func three() float32 {
	threeOnce.Do(func() { // 使用Once初始化
		fmt.Println("threeOnce init!")
		threeOnce.v = 3.0
	})
	return threeOnce.v
}

func TestOnce(t *testing.T) {
	t.Log(three())
	t.Log(three())
}
