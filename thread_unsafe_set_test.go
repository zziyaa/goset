package goset

import "testing"

func TestNewSet(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		s := NewSet[int](true)
		if s.Size() != 0 {
			t.Errorf("expected empty set, got size %d", s.Size())
		}
		if !s.IsEmpty() {
			t.Errorf("expected empty set to be empty")
		}
	})

	t.Run("with items", func(t *testing.T) {
		items := []int{1, 2, 3}
		s := NewSet(true, items...)
		if s.Size() != len(items) {
			t.Errorf("expected size %d, got %d", len(items), s.Size())
		}
		for _, item := range items {
			if !s.Contains(item) {
				t.Errorf("expected set to contain %d", item)
			}
		}
	})

	t.Run("with duplicate items", func(t *testing.T) {
		s := NewSet(true, 1, 2, 2, 3, 3, 3)
		if s.Size() != 3 {
			t.Errorf("expected size 3, got %d", s.Size())
		}
		for _, item := range []int{1, 2, 3} {
			if !s.Contains(item) {
				t.Errorf("expected set to contain %d", item)
			}
		}
	})
}

func TestSet_Contains(t *testing.T) {
	t.Run("existing item", func(t *testing.T) {
		s := NewSet(true, 1, 2, 3)
		if !s.Contains(2) {
			t.Errorf("expected set to contain 2")
		}
	})

	t.Run("non-existing item", func(t *testing.T) {
		s := NewSet(true, 1, 2, 3)
		if s.Contains(4) {
			t.Errorf("expected set to not contain 4")
		}
	})

	t.Run("empty set", func(t *testing.T) {
		s := NewSet[int](true)
		if s.Contains(1) {
			t.Errorf("expected empty set to not contain any items")
		}
	})
}

func TestSet_Insert(t *testing.T) {
	t.Run("insert new item", func(t *testing.T) {
		s := NewSet[int](true)
		s.Insert(1)
		if !s.Contains(1) {
			t.Errorf("expected set to contain 1 after insertion")
		}
		if s.Size() != 1 {
			t.Errorf("expected size 1, got %d", s.Size())
		}
	})

	t.Run("insert existing item", func(t *testing.T) {
		s := NewSet(true, 1, 2, 3)
		initialSize := s.Size()
		s.Insert(2)
		if s.Size() != initialSize {
			t.Errorf("expected size to remain %d, got %d", initialSize, s.Size())
		}
		if !s.Contains(2) {
			t.Errorf("expected set to still contain 2")
		}
	})

	t.Run("multiple inserts", func(t *testing.T) {
		s := NewSet[int](true)
		items := []int{1, 2, 3, 4, 5}
		for _, item := range items {
			s.Insert(item)
		}
		if s.Size() != len(items) {
			t.Errorf("expected size %d, got %d", len(items), s.Size())
		}
		for _, item := range items {
			if !s.Contains(item) {
				t.Errorf("expected set to contain %d", item)
			}
		}
	})
}

func TestSet_Remove(t *testing.T) {
	t.Run("remove existing item", func(t *testing.T) {
		items := []int{3, 4, 5, 6, 7}
		s := NewSet(true, items...)

		// Ensure the set contains some items
		if s.Size() != len(items) {
			t.Fatalf("expected size: %d, actual size: %d", len(items), s.Size())
		}

		i := 1 // 0 <= i < len(items)

		if !s.Contains(items[i]) {
			t.Fatalf("set should have contain the item: %d", items[i])
		}

		s.Remove(items[i])

		if s.Size() != len(items)-1 {
			t.Fatalf("expected size: %d, actual size: %d", len(items), s.Size())
		}

		if s.Contains(items[i]) {
			t.Fatalf("set should NOT have contain the item: %d", items[i])
		}
	})

	t.Run("try to remove a non-existing item", func(t *testing.T) {
		items := []int{3, 4, 5, 6, 7}
		s := NewSet(true, items...)

		// Ensure the set contains some items
		if s.Size() != len(items) {
			t.Fatalf("expected size: %d, actual size: %d", len(items), s.Size())
		}

		val := 1022 // random number outside of the `items` values

		if s.Contains(val) {
			t.Fatalf("set should NOT have contain the item: %d", val)
		}

		s.Remove(val)

		if s.Size() != len(items) {
			t.Fatalf("expected size: %d, actual size: %d", len(items), s.Size())
		}

		if s.Contains(val) {
			t.Fatalf("set should NOT have contain the item: %d", val)
		}
	})

}

func TestSet_Size(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		s := NewSet[int](true)
		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})

	t.Run("with some items", func(t *testing.T) {
		items := []int{3, 4, 5}
		s := NewSet(true, items...)
		if s.Size() != len(items) {
			t.Fatalf("expected size: %d, actual size: %d", len(items), s.Size())
		}
	})

	t.Run("after modifications", func(t *testing.T) {
		s := NewSet(true, 1, 2, 3)
		if s.Size() != 3 {
			t.Errorf("expected initial size 3, got %d", s.Size())
		}
		s.Insert(4)
		if s.Size() != 4 {
			t.Errorf("expected size 4 after insert, got %d", s.Size())
		}
		s.Remove(1)
		if s.Size() != 3 {
			t.Errorf("expected size 3 after remove, got %d", s.Size())
		}
	})
}

func TestSet_IsEmpty(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		s := NewSet[int](true)
		if !s.IsEmpty() {
			t.Errorf("set should be empty, but it's not")
		}
	})

	t.Run("non-empty set", func(t *testing.T) {
		s := NewSet(true, 3, 4, 5)
		if s.IsEmpty() {
			t.Errorf("set should be non-empty, but it's not")
		}
	})
}

func TestSet_Clear(t *testing.T) {
	t.Run("clear empty set", func(t *testing.T) {
		s := NewSet[int](true)
		s.Clear()
		if !s.IsEmpty() {
			t.Errorf("expected set to remain empty after clear")
		}
		if s.Size() != 0 {
			t.Errorf("expected size 0 after clear, got %d", s.Size())
		}
	})

	t.Run("clear non-empty set", func(t *testing.T) {
		items := []int{3, 4, 5}
		s := NewSet(true, items...)
		if s.IsEmpty() || s.Size() != len(items) {
			t.Fatalf("expected size: %d, actual size: %d", len(items), s.Size())
		}
		s.Clear()
		if !s.IsEmpty() || s.Size() != 0 {
			t.Errorf("set shouldn't contain any items, but it does")
		}
		for _, item := range items {
			if s.Contains(item) {
				t.Errorf("expected item %d to be removed after clear", item)
			}
		}
	})
}

func TestSet_ToSlice(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		s := NewSet[int](true)
		slice := s.ToSlice()
		if len(slice) != 0 {
			t.Errorf("expected empty slice, got length %d", len(slice))
		}
	})

	t.Run("single item", func(t *testing.T) {
		s := NewSet(true, 42)
		slice := s.ToSlice()
		if len(slice) != 1 {
			t.Errorf("expected slice length 1, got %d", len(slice))
		}
		if slice[0] != 42 {
			t.Errorf("expected slice to contain 42, got %d", slice[0])
		}
	})

	t.Run("multiple items", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		s := NewSet(true, items...)
		slice := s.ToSlice()

		if len(slice) != len(items) {
			t.Errorf("expected slice length %d, got %d", len(items), len(slice))
		}

		// Create a map to check all items are present (order doesn't matter)
		found := make(map[int]bool)
		for _, item := range slice {
			found[item] = true
		}

		for _, item := range items {
			if !found[item] {
				t.Errorf("expected slice to contain %d", item)
			}
		}
	})

	t.Run("after modifications", func(t *testing.T) {
		s := NewSet(true, 1, 2, 3)
		s.Insert(4)
		s.Remove(2)

		slice := s.ToSlice()
		if len(slice) != 3 {
			t.Errorf("expected slice length 3, got %d", len(slice))
		}

		// Check expected items are present
		found := make(map[int]bool)
		for _, item := range slice {
			found[item] = true
		}

		expectedItems := []int{1, 3, 4}
		for _, item := range expectedItems {
			if !found[item] {
				t.Errorf("expected slice to contain %d", item)
			}
		}

		// Check removed item is not present
		if found[2] {
			t.Errorf("expected slice to not contain removed item 2")
		}
	})

	t.Run("slice independence", func(t *testing.T) {
		s := NewSet(true, 1, 2, 3)
		slice := s.ToSlice()

		// Modify the returned slice
		if len(slice) > 0 {
			slice[0] = 999
		}

		// Original set should be unchanged
		if s.Contains(999) {
			t.Errorf("modifying returned slice should not affect original set")
		}
		if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
			t.Errorf("original set should remain unchanged")
		}
	})
}
