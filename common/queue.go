package common

// Queue is a simple FIFO queue.
type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Pop() T {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}
