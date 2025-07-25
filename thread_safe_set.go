package goset

import "sync"

type ThreadSafeSet[T comparable] struct {
	set *ThreadUnsafeSet[T]
	mu  sync.RWMutex
}

// newThreadSafeSet initializes and returns a new thread-safe set with the given items
func newThreadSafeSet[T comparable](items ...T) *ThreadSafeSet[T] {
	return &ThreadSafeSet[T]{
		set: newThreadUnsafeSet(items...),
	}
}

func (s *ThreadSafeSet[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.set.Contains(item)
}

func (s *ThreadSafeSet[T]) Insert(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.set.Insert(item)
}

func (s *ThreadSafeSet[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.set.Remove(item)
}

func (s *ThreadSafeSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.set.Size()
}

func (s *ThreadSafeSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.set.IsEmpty()
}

func (s *ThreadSafeSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.set.Clear()
}

func (s *ThreadSafeSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.set.ToSlice()
}
