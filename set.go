package goset

// Simple, generic, non-thread safe set data structure that provides a conventional API
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet initializes and returns a new set with the given items
func NewSet[T comparable](items ...T) *Set[T] {
	data := make(map[T]struct{}, len(items))
	for _, item := range items {
		data[item] = struct{}{}
	}

	return &Set[T]{
		data: data,
	}
}

// Contains returns true if the item exists
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.data[item]
	return ok
}

// Insert adds the item into the set
func (s *Set[T]) Insert(item T) {
	s.data[item] = struct{}{}
}

// Remove deletes the item from the set,
// it does nothing if item doesn't exist in the set
func (s *Set[T]) Remove(item T) {
	delete(s.data, item)
}

// Size returns the count of items in the set
func (s *Set[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the set is empty
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Clear removes all items from the set
func (s *Set[T]) Clear() {
	clear(s.data)
}

// ToSlice returns all the items in the set as a slice. Note that the returned slice is unordered.
func (s *Set[T]) ToSlice() []T {
	items := make([]T, 0, s.Size())
	for item := range s.data {
		items = append(items, item)
	}
	return items
}
