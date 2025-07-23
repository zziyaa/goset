package goset

type Set[T comparable] interface {
	// Contains returns true if the item exists
	Contains(item T) bool

	// Insert adds the item into the set
	Insert(item T)

	// Remove deletes the item from the set,
	// it does nothing if item doesn't exist in the set
	Remove(item T)

	// Size returns cardinality of the set
	Size() int

	// IsEmpty returns true if the set is empty
	IsEmpty() bool

	// Clear removes all items from the set
	Clear()

	// ToSlice returns all the items in the set as a slice. Note that the returned slice is unordered
	ToSlice() []T
}

func NewSet[T comparable](singleThread bool, items ...T) Set[T] {
	if singleThread {
		return newThreadUnsafeSet(items...)
	}

	// TODO: Return thread safe set.
	return nil
}
