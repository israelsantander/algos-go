package sorting

import (
	"cmp"
	"slices"
)

// Heap returns a sorted copy of values using heap sort.
//
// # Characteristics
//
// Heap sort guarantees O(n log n) time and uses only constant extra working memory in its
// in-place form. It is a good fit when worst-case time matters and stability does not.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(n) additional space for the copied result.
func Heap[T cmp.Ordered](values []T) []T {
	return HeapFunc(values, cmp.Less[T])
}

// HeapFunc returns a sorted copy of values using heap sort and the provided comparator.
//
// # Use
//
// Use it for custom types when you want deterministic O(n log n) sorting without merge sort's extra buffer logic.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(n) additional space for the copied result.
func HeapFunc[T any](values []T, less func(a, b T) bool) []T {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	out := slices.Clone(values)
	heapInPlace(out, less)
	return out
}

// HeapInPlace sorts values in place using heap sort.
//
// # Characteristics
//
// It is in-place and guarantees O(n log n) time, but it is not stable and often has larger
// constant factors than quicksort in practice.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(1) additional space.
func HeapInPlace[T cmp.Ordered](values []T) {
	HeapInPlaceFunc(values, cmp.Less[T])
}

// HeapInPlaceFunc sorts values in place using heap sort and the provided comparator.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(1) additional space.
func HeapInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	heapInPlace(values, less)
}

// heapInPlace first builds a max-heap, then repeatedly moves the largest element to the end
// of the slice and shrinks the heap boundary.
func heapInPlace[T any](values []T, less func(a, b T) bool) {
	n := len(values)
	for root := n/2 - 1; root >= 0; root-- {
		siftDown(values, root, n, less)
	}

	for end := n - 1; end > 0; end-- {
		values[0], values[end] = values[end], values[0]
		siftDown(values, 0, end, less)
	}
}

// siftDown restores the max-heap property by pushing one out-of-place root downward until
// it dominates both of its children again.
func siftDown[T any](values []T, root, size int, less func(a, b T) bool) {
	for {
		left := 2*root + 1
		if left >= size {
			return
		}

		largest := left
		right := left + 1
		if right < size && less(values[largest], values[right]) {
			largest = right
		}
		if !less(values[root], values[largest]) {
			return
		}

		values[root], values[largest] = values[largest], values[root]
		root = largest
	}
}
