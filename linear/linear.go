package linear

// Stack is a generic last-in, first-out container backed by a slice.
type Stack[T any] struct {
	items []T
}

// Push adds value to the top of the stack.
// Time complexity: O(1) amortized.
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the top stack value.
// It reports false when the stack is empty.
// Time complexity: O(1).
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := len(s.items) - 1
	value := s.items[last]
	var cleared T
	s.items[last] = cleared
	s.items = s.items[:last]
	return value, true
}

// Peek returns the top stack value without removing it.
// It reports false when the stack is empty.
// Time complexity: O(1).
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Len reports the number of values currently stored in the stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}

// IsEmpty reports whether the stack currently contains no values.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Values returns a copy of the stack contents in bottom-to-top order.
func (s *Stack[T]) Values() []T {
	return append([]T(nil), s.items...)
}

// Queue is a generic first-in, first-out container backed by a slice and a moving head index.
type Queue[T any] struct {
	items []T
	head  int
}

// Enqueue adds value to the back of the queue.
// Time complexity: O(1) amortized.
func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

// Dequeue removes and returns the front queue value.
// It reports false when the queue is empty.
// Time complexity: O(1) amortized.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}

	value := q.items[q.head]
	var cleared T
	q.items[q.head] = cleared
	q.head++

	switch {
	case q.head == len(q.items):
		q.items = nil
		q.head = 0
	case q.head >= len(q.items)/2:
		q.compact()
	}

	return value, true
}

// Peek returns the front queue value without removing it.
// It reports false when the queue is empty.
// Time complexity: O(1).
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.IsEmpty() {
		return zero, false
	}
	return q.items[q.head], true
}

// Len reports the number of values currently stored in the queue.
func (q *Queue[T]) Len() int {
	return len(q.items) - q.head
}

// IsEmpty reports whether the queue currently contains no values.
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

// Values returns a copy of the queue contents in front-to-back order.
func (q *Queue[T]) Values() []T {
	return append([]T(nil), q.items[q.head:]...)
}

// compact discards consumed prefix capacity once it becomes a meaningful share of the backing slice.
func (q *Queue[T]) compact() {
	q.items = append([]T(nil), q.items[q.head:]...)
	q.head = 0
}
