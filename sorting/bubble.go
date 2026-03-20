package sorting

import (
	"cmp"
	"slices"
)

// Bubble returns a sorted copy of values using bubble sort.
//
// It is a stable sort, meaning it preserves the original order of equal elements.
// Time complexity: O(n^2) average and worst case, O(n) best case.
// Space complexity: O(n) for the returned copy.
func Bubble[T cmp.Ordered](values []T) []T {
	return BubbleFunc(values, cmp.Less[T])
}

// BubbleFunc returns a sorted copy of values using bubble sort and the provided comparator.
func BubbleFunc[T any](values []T, less func(a, b T) bool) []T {
	out := slices.Clone(values)
	BubbleInPlaceFunc(out, less)
	return out
}

// BubbleInPlace sorts values in place in ascending order using bubble sort.
//
// Time complexity: O(n^2) average and worst case, O(n) best case.
// Space complexity: O(1).
func BubbleInPlace[T cmp.Ordered](values []T) {
	BubbleInPlaceFunc(values, cmp.Less[T])
}

// BubbleInPlaceFunc sorts values in place using bubble sort and the provided comparator.
// It panics if less is nil.
func BubbleInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	bubbleInPlace(values, less)
}

// bubbleInPlace tracks the last swap position so each pass only revisits the unsorted suffix.
// This allows bubble sort to finish in O(n) time for already-sorted input.
func bubbleInPlace[T any](values []T, less func(a, b T) bool) {
	for n := len(values); n > 1; {
		lastSwap := 0
		for i := 1; i < n; i++ {
			if less(values[i], values[i-1]) {
				values[i-1], values[i] = values[i], values[i-1]
				lastSwap = i
			}
		}
		if lastSwap == 0 {
			return
		}
		n = lastSwap
	}
}
