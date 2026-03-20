package sorting

import "cmp"

// QuickSelect returns the k-th smallest value from a copy of values using quickselect.
//
// It is useful when you need one order statistic, such as the median, without fully sorting the
// slice. The function returns false when k is out of range.
// Example: QuickSelect([]int{9, 1, 7, 3, 5}, 2) returns 5, true.
// Average time complexity: O(n). Worst case: O(n^2). Additional space: O(n) for the copied input.
func QuickSelect[T cmp.Ordered](values []T, k int) (T, bool) {
	return QuickSelectFunc(values, k, orderedLess[T])
}

// QuickSelectFunc returns the k-th smallest value from a copy of values using quickselect and less.
//
// The function returns false when k is out of range.
// Example: QuickSelectFunc(players, 4, func(a, b player) bool { return a.Score < b.Score }).
// Average time complexity: O(n). Worst case: O(n^2). Additional space: O(n) for the copied input.
func QuickSelectFunc[T any](values []T, k int, less func(a, b T) bool) (T, bool) {
	var zero T
	if k < 0 || k >= len(values) {
		return zero, false
	}
	out := clone(values)
	return quickSelect(out, k, requireLess(less)), true
}

func quickSelect[T any](values []T, k int, less func(a, b T) bool) T {
	low, high := 0, len(values)-1
	for {
		if high-low+1 <= quickInsertionCutoff {
			insertionInPlace(values[low:high+1], less)
			return values[k]
		}

		pivot := medianOfThree(values, low, high, less)
		lt, gt := partition(values, low, high, pivot, less)
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
