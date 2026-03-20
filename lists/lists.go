package lists

// SinglyNode is a node in a singly linked list.
type SinglyNode[T any] struct {
	Value T
	Next  *SinglyNode[T]
}

// SinglyLinkedList is a generic singly linked list with cached head, tail, and length.
type SinglyLinkedList[T any] struct {
	head *SinglyNode[T]
	tail *SinglyNode[T]
	len  int
}

// Append adds value to the end of the list.
// Time complexity: O(1).
func (l *SinglyLinkedList[T]) Append(value T) {
	node := &SinglyNode[T]{Value: value}
	if l.tail == nil {
		l.head, l.tail = node, node
	} else {
		l.tail.Next = node
		l.tail = node
	}
	l.len++
}

// Prepend adds value to the front of the list.
// Time complexity: O(1).
func (l *SinglyLinkedList[T]) Prepend(value T) {
	node := &SinglyNode[T]{Value: value, Next: l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.len++
}

// InsertAt inserts value at index and reports whether the insertion succeeded.
// Valid indexes are in the inclusive range [0, Len()].
// Time complexity: O(n).
func (l *SinglyLinkedList[T]) InsertAt(index int, value T) bool {
	if index < 0 || index > l.len {
		return false
	}
	if index == 0 {
		l.Prepend(value)
		return true
	}
	if index == l.len {
		l.Append(value)
		return true
	}
	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	node := &SinglyNode[T]{Value: value, Next: prev.Next}
	prev.Next = node
	l.len++
	return true
}

// DeleteAt removes and returns the value at index.
// It reports false when index is out of range.
// Time complexity: O(n).
func (l *SinglyLinkedList[T]) DeleteAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.len || l.head == nil {
		return zero, false
	}
	if index == 0 {
		value := l.head.Value
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
		l.len--
		return value, true
	}
	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	target := prev.Next
	prev.Next = target.Next
	if target == l.tail {
		l.tail = prev
	}
	l.len--
	return target.Value, true
}

// Len reports the number of values currently stored in the list.
func (l *SinglyLinkedList[T]) Len() int { return l.len }

// Values returns the list contents from head to tail as a slice copy.
func (l *SinglyLinkedList[T]) Values() []T {
	out := make([]T, 0, l.len)
	for node := l.head; node != nil; node = node.Next {
		out = append(out, node.Value)
	}
	return out
}

// DoublyNode is a node in a doubly linked list.
type DoublyNode[T any] struct {
	Value T
	Prev  *DoublyNode[T]
	Next  *DoublyNode[T]
}

// DoublyLinkedList is a generic doubly linked list with cached head, tail, and length.
type DoublyLinkedList[T any] struct {
	head *DoublyNode[T]
	tail *DoublyNode[T]
	len  int
}

// Append adds value to the end of the list.
// Time complexity: O(1).
func (l *DoublyLinkedList[T]) Append(value T) {
	node := &DoublyNode[T]{Value: value, Prev: l.tail}
	if l.tail != nil {
		l.tail.Next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.len++
}

// Prepend adds value to the front of the list.
// Time complexity: O(1).
func (l *DoublyLinkedList[T]) Prepend(value T) {
	node := &DoublyNode[T]{Value: value, Next: l.head}
	if l.head != nil {
		l.head.Prev = node
	} else {
		l.tail = node
	}
	l.head = node
	l.len++
}

// DeleteAt removes and returns the value at index.
// It reports false when index is out of range.
// Time complexity: O(min(index, Len()-1-index)).
func (l *DoublyLinkedList[T]) DeleteAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.len {
		return zero, false
	}
	curr := l.nodeAt(index)
	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	} else {
		l.head = curr.Next
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	} else {
		l.tail = curr.Prev
	}
	l.len--
	return curr.Value, true
}

// nodeAt walks from the closer end of the list to reach index with fewer pointer hops.
func (l *DoublyLinkedList[T]) nodeAt(index int) *DoublyNode[T] {
	if index < l.len/2 {
		curr := l.head
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		return curr
	}

	curr := l.tail
	for i := l.len - 1; i > index; i-- {
		curr = curr.Prev
	}
	return curr
}

// Len reports the number of values currently stored in the list.
func (l *DoublyLinkedList[T]) Len() int { return l.len }

// Values returns the list contents from head to tail as a slice copy.
func (l *DoublyLinkedList[T]) Values() []T {
	out := make([]T, 0, l.len)
	for node := l.head; node != nil; node = node.Next {
		out = append(out, node.Value)
	}
	return out
}

// CircularLinkedList is a generic circular singly linked list.
type CircularLinkedList[T any] struct {
	head *SinglyNode[T]
	tail *SinglyNode[T]
	len  int
}

// Append adds value to the end of the circular list.
// Time complexity: O(1).
func (l *CircularLinkedList[T]) Append(value T) {
	node := &SinglyNode[T]{Value: value}
	if l.head == nil {
		node.Next = node
		l.head, l.tail = node, node
	} else {
		node.Next = l.head
		l.tail.Next = node
		l.tail = node
	}
	l.len++
}

// Values returns up to limit values by walking from the head around the cycle once.
func (l *CircularLinkedList[T]) Values(limit int) []T {
	if l.head == nil || limit <= 0 {
		return nil
	}
	if limit > l.len {
		limit = l.len
	}
	out := make([]T, 0, limit)
	curr := l.head
	for i := 0; i < limit; i++ {
		out = append(out, curr.Value)
		curr = curr.Next
	}
	return out
}

// Len reports the number of values currently stored in the list.
func (l *CircularLinkedList[T]) Len() int { return l.len }

// CircularDoublyNode is a node in a circular doubly linked list.
type CircularDoublyNode[T any] struct {
	Value T
	Prev  *CircularDoublyNode[T]
	Next  *CircularDoublyNode[T]
}

// CircularDoublyLinkedList is a generic circular doubly linked list.
type CircularDoublyLinkedList[T any] struct {
	head *CircularDoublyNode[T]
	len  int
}

// Append adds value to the end of the circular doubly linked list.
// Time complexity: O(1).
func (l *CircularDoublyLinkedList[T]) Append(value T) {
	node := &CircularDoublyNode[T]{Value: value}
	if l.head == nil {
		node.Next, node.Prev = node, node
		l.head = node
		l.len = 1
		return
	}
	tail := l.head.Prev
	node.Next = l.head
	node.Prev = tail
	tail.Next = node
	l.head.Prev = node
	l.len++
}

// Values returns up to limit values by walking from the head around the cycle once.
func (l *CircularDoublyLinkedList[T]) Values(limit int) []T {
	if l.head == nil || limit <= 0 {
		return nil
	}
	if limit > l.len {
		limit = l.len
	}
	out := make([]T, 0, limit)
	curr := l.head
	for i := 0; i < limit; i++ {
		out = append(out, curr.Value)
		curr = curr.Next
	}
	return out
}

// Len reports the number of values currently stored in the list.
func (l *CircularDoublyLinkedList[T]) Len() int { return l.len }
