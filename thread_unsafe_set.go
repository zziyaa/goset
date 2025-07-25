package goset

type ThreadUnsafeSet[T comparable] struct {
	data map[T]struct{}
}

// newThreadUnsafeSet initializes and returns a new thread-unsafe set with the given items
func newThreadUnsafeSet[T comparable](items ...T) *ThreadUnsafeSet[T] {
	data := make(map[T]struct{}, len(items))
	for _, item := range items {
		data[item] = struct{}{}
	}

	return &ThreadUnsafeSet[T]{
		data: data,
	}
}

func (s *ThreadUnsafeSet[T]) Contains(item T) bool {
	_, ok := s.data[item]
	return ok
}

func (s *ThreadUnsafeSet[T]) Insert(item T) {
	s.data[item] = struct{}{}
}

func (s *ThreadUnsafeSet[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *ThreadUnsafeSet[T]) Size() int {
	return len(s.data)
}

func (s *ThreadUnsafeSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *ThreadUnsafeSet[T]) Clear() {
	clear(s.data)
}

func (s *ThreadUnsafeSet[T]) ToSlice() []T {
	items := make([]T, 0, s.Size())
	for item := range s.data {
		items = append(items, item)
	}

	return items
}
