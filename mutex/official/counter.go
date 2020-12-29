package official

import "sync"

type counter struct {
	sync.Mutex
	count int
}

func (c *counter) Incr(delta int) {
	c.Lock()
	defer c.Unlock()
	c.count += delta
}

func (c *counter) Count() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}
