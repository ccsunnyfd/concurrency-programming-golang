package oncehasdonemethod

import (
	"testing"
	"time"
)

func TestOnceHasDoneMethod(t *testing.T) {
	var flag Once
	t.Log(flag.Done()) //false
	flag.Do(func() {
		time.Sleep(time.Second)
	})
	t.Log(flag.Done()) //true
}
