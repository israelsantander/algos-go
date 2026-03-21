package sorting

import (
	"cmp"
	"slices"
)

// QuickSelect returns the k-th smallest value from a copy of values using quickselect.
//
// # Behavior
//
// Use it when you need one order statistic, such as the median, without fully sorting the slice.
//
// It returns false when k is out of range.
//
// # Complexity
//
// It runs in O(n) time on average, O(n^2) in the worst case, and uses O(n) additional space.
func QuickSelect[T cmp.Ordered](values []T, k int) (T, bool) {
	return QuickSelectFunc(values, k, cmp.Less[T])
}

// QuickSelectFunc returns the k-th smallest value from a copy of values using quickselect and the provided comparator.
//
// # Behavior
//
// It returns false when k is out of range.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n) time on average, O(n^2) in the worst case, and uses O(n) additional space.
func QuickSelectFunc[T any](values []T, k int, less func(a, b T) bool) (T, bool) {
	var zero T
	if k < 0 || k >= len(values) {
		return zero, false
	}
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	out := slices.Clone(values)
	return quickSelect(out, k, less), true
}

func quickSelect[T any](values []T, k int, less func(a, b T) bool) T {
	low, high := 0, len(values)-1
	for {
		if high-low+1 <= quickInsertionCutoff {
			insertionInPlace(values[low:high+1], less)
			return values[k]
		}

		pivot := medianOfThreeValue(values, low, high, less)
		lt, gt := partition3Way(values, low, high, pivot, less)
		switch {
		case k < lt:
			high = lt - 1
		case k > gt:
			low = gt + 1
		default:
			return values[k]
		}
	}
}
