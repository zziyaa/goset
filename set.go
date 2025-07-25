package goset

type Set[T comparable] interface {
	// Contains returns true if the item exists
	Contains(item T) bool

	// Insert adds the item into the set
	Insert(item T)

	// Remove deletes the item from the set,
	// it does nothing if the item doesn't exist in the set
	Remove(item T)

	// Size returns size of the set
	Size() int

	// IsEmpty returns true if the set is empty
	IsEmpty() bool

	// Clear removes all items from the set
	Clear()

	// ToSlice returns all the items in the set as a slice. Note that the returned slice is unordered
	ToSlice() []T
}

// NewThreadUnsafeSet creates a set for single threaded use with initial items (if any)
func NewThreadUnsafeSet[T comparable](items ...T) Set[T] {
	return newThreadUnsafeSet(items...)
}

// NewThreadSafeSet creates a set for concurrent use with initial items (if any)
func NewThreadSafeSet[T comparable](items ...T) Set[T] {
	return newThreadSafeSet(items...)
}
