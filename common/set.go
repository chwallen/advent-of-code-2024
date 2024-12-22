package common

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

// Clear turns the set s into the empty set.
func (s Set[T]) Clear() {
	clear(s.items)
}

// Contains checks if the item v is present in the set s.
func (s Set[T]) Contains(v T) bool {
	_, found := s.items[v]
	return found
}

// Len gets the number of items in the set s.
func (s Set[T]) Len() int {
	return len(s.items)
}
