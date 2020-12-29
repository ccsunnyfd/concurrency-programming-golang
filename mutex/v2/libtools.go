package v2

type semaphore struct {
	ch chan struct{}
}

func newSemaphore() *semaphore {
	return &semaphore{
		ch: make(chan struct{}), // chan length must be 0, otherwise goroutine won't block correctly
	}
}

// semacquire is a mocked function, should be C primitive implementation in reality
func (s *semaphore) semacquire() {
	s.ch <- struct{}{}
}

// semarelease is a mocked function, should be C primitive implementation in reality
func (s *semaphore) semarelease() {
	<-s.ch
}

func (s *semaphore) close() {
	close(s.ch)
}
