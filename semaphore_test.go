package semaphore

import (
	"testing"
	"time"
)

func TestNewSemaphore(t *testing.T) {
	s := NewSemaphore(5)
	if s.capacity != 5 {
		t.Errorf("Semaphore capacity expected to be %d, got %d", 5, s.capacity)
	}
	if s.count != 0 {
		t.Errorf("Semaphore count expected to be %d, got %d", 0, s.count)
	}
}

func TestAcquireAndRelease(t *testing.T) {
	s := NewSemaphore(2)
	s.Acquire()
	s.Acquire()
	if s.count != 2 {
		t.Errorf("Semaphore count expected to be %d, got %d", 2, s.count)
	}
	done := make(chan bool)
	go func() {
		s.Release()
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Errorf("Semaphore release took too long")
	}
	if s.count != 1 {
		t.Errorf("Semaphore count expected to be %d, got %d", 1, s.count)
	}
}

func TestTryAcquire(t *testing.T) {
	s := NewSemaphore(1)
	if !s.TryAcquire() {
		t.Errorf("Semaphore should have been acquired successfully")
	}
	if s.TryAcquire() {
		t.Errorf("Semaphore should not have been acquired successfully")
	}
	if s.count != 1 {
		t.Errorf("Semaphore count expected to be %d, got %d", 1, s.count)
	}
}

func TestAvailable(t *testing.T) {
	s := NewSemaphore(3)
	s.Acquire()
	s.Acquire()
	if s.Available() != 1 {
		t.Errorf("Semaphore available permits expected to be %d, got %d", 1, s.Available())
	}
	s.Acquire()
	if s.Available() != 0 {
		t.Errorf("Semaphore available permits expected to be %d, got %d", 0, s.Available())
	}
}
