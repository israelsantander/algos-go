package sorting

import (
	"cmp"
	"slices"
)

// Selection returns a sorted copy of values using selection sort.
//
// # Characteristics
//
// Selection sort performs a fixed number of comparisons regardless of input order.
//
// It is simple and predictable, but usually slower than insertion sort on small arrays.
//
// # Complexity
//
// It runs in O(n^2) time and uses O(n) additional space for the copied result.
func Selection[T cmp.Ordered](values []T) []T {
	return SelectionFunc(values, cmp.Less[T])
}

// SelectionFunc returns a sorted copy of values using selection sort and the provided comparator.
//
// # Use
//
// Use it when you want the selection-sort behavior for custom types.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// # Complexity
//
// It runs in O(n^2) time and uses O(n) additional space for the copied result.
func SelectionFunc[T any](values []T, less func(a, b T) bool) []T {
	out := slices.Clone(values)
	SelectionInPlaceFunc(out, less)
	return out
}

// SelectionInPlace sorts values in place using selection sort.
//
// # Characteristics
//
// Selection sort minimizes swaps, which can help when writes are expensive, but it is not stable.
//
// # Complexity
//
// It runs in O(n^2) time and uses O(1) additional space.
func SelectionInPlace[T cmp.Ordered](values []T) {
	SelectionInPlaceFunc(values, cmp.Less[T])
}

// SelectionInPlaceFunc sorts values in place using selection sort and the provided comparator.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n^2) time and uses O(1) additional space.
func SelectionInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	selectionInPlace(values, less)
}

// selectionInPlace shrinks the unsorted window from both ends by placing its minimum and maximum
// element on each pass.
func selectionInPlace[T any](values []T, less func(a, b T) bool) {
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		minIndex, maxIndex := left, left
		for i := left + 1; i <= right; i++ {
			if less(values[i], values[minIndex]) {
				minIndex = i
			}
			if less(values[maxIndex], values[i]) {
				maxIndex = i
			}
		}
		if minIndex != left {
			values[left], values[minIndex] = values[minIndex], values[left]
			if maxIndex == left {
				maxIndex = minIndex
			}
		}
		if maxIndex != right {
			values[right], values[maxIndex] = values[maxIndex], values[right]
		}
	}
}
