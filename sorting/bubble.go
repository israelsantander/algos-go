package sorting

import (
	"cmp"
	"slices"
)

// Bubble returns a sorted copy of values using bubble sort.
//
// # Characteristics
//
// Bubble sort is stable, so equal elements keep their original order.
//
// # Complexity
//
// It runs in O(n^2) time in the average and worst cases, O(n) on already-sorted input,
// and uses O(n) additional space for the copied result.
func Bubble[T cmp.Ordered](values []T) []T {
	return BubbleFunc(values, cmp.Less[T])
}

// BubbleFunc returns a sorted copy of values using bubble sort and less.
//
// # Use
//
// Use it for custom element types or custom orderings.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
func BubbleFunc[T any](values []T, less func(a, b T) bool) []T {
	out := slices.Clone(values)
	BubbleInPlaceFunc(out, less)
	return out
}

// BubbleInPlace sorts values in place using bubble sort.
//
// # Characteristics
//
// Bubble sort is stable.
//
// # Complexity
//
// It runs in O(n^2) time in the average and worst cases, O(n) on already-sorted input,
// and uses O(1) additional space.
func BubbleInPlace[T cmp.Ordered](values []T) {
	BubbleInPlaceFunc(values, cmp.Less[T])
}

// BubbleInPlaceFunc sorts values in place using bubble sort and less.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
func BubbleInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	bubbleInPlace(values, less)
}

// bubbleInPlace shortens each pass to the last swap position.
// That lets it stop early for already-sorted input.
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
