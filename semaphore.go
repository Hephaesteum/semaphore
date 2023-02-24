package semaphore

import "sync"

type Semaphore struct {
	count    int
	capacity int
	mutex    sync.Mutex
	cond     *sync.Cond
}

// Return new semaphore with specified capacity.
func NewSemaphore(capacity int) *Semaphore {
	s := &Semaphore{
		capacity: capacity,
		cond:     sync.NewCond(&sync.Mutex{}),
	}
	return s
}

// Acquire acquires the semaphore, blocking until it becomes available if necessary.
func (s *Semaphore) Acquire() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for s.count == s.capacity {
		s.cond.Wait()
	}
	s.count++
}

// Release releases the semaphore, allowing another go routine to acquire it.
func (s *Semaphore) Release() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count--
	s.cond.Signal()
}

// TryAcquire attempts to acquire the semaphore without blocking, returning
//  true if successful.
func (s *Semaphore) TryAcquire() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.count == s.capacity {
		return false
	}
	s.count++
	return true
}

// Available returns number of permits available in the semaphore.
func (s *Semaphore) Available() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.capacity - s.count
}
