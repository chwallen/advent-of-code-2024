package common

import (
	"iter"
	"maps"
	"slices"
)

type empty struct{}

type Set[T comparable] struct {
	items map[T]empty
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{make(map[T]empty)}
}

// Add tries to add v to the set s. Returns false if v is already in the set.
func (s Set[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}
	s.items[v] = empty{}
	return true
}

// All returns an iterator over the values in the set s. The iteration order is
// unspecified and may differ between invokations.
func (s Set[T]) All() iter.Seq[T] {
	return maps.Keys(s.items)
}

// Clear turns the set s into the empty set.
func (s Set[T]) Clear() {
	clear(s.items)
}

// Clone creates a clone of set s with shallowly copied values.
func (s Set[T]) Clone() Set[T] {
	return Set[T]{items: maps.Clone(s.items)}
}

// Contains checks if the item v is present in the set s.
func (s Set[T]) Contains(v T) bool {
	_, found := s.items[v]
	return found
}

// Delete removes an item from the set s if it exists.
func (s Set[T]) Delete(item T) {
	delete(s.items, item)
}

// Difference creates a new set of all items in s which are not present in the
// other set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := maps.Clone(s.items)
	for item := range other.items {
		if _, contains := diff[item]; contains {
			delete(diff, item)
		}
	}
	return Set[T]{diff}
}

// Intersection creates a new set of the items present in both set s and other.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	items := make(map[T]empty)
	for item := range s.items {
		if other.Contains(item) {
			items[item] = empty{}
		}
	}
	return Set[T]{items}
}

// Items returns an unordered slice of all items in the set s.
func (s Set[T]) Items() []T {
	return slices.Collect(s.All())
}

// Len gets the number of items in the set s.
func (s Set[T]) Len() int {
	return len(s.items)
}

// Union creates a new set of all items present in either set s or other.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := maps.Clone(s.items)
	maps.Copy(union, other.items)
	return Set[T]{union}
}
