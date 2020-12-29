package v1

type counter struct {
	mu    Mutex
	count int
}

func newCounter() *counter {
	return &counter{
		mu:    newMutex(),
		count: 0,
	}
}

func (c *counter) Incr(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count += delta
}

func (c *counter) Count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *counter) Close() {
	c.mu.Close()
}
