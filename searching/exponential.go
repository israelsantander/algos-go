package searching

import "cmp"

// Exponential searches a sorted slice by first growing a probe window exponentially and then
// using binary search inside that window.
//
// Exponential search is useful when the target is likely to appear near the front of a very large
// sorted slice. For duplicate values, Exponential returns the leftmost matching index.
// Example: Exponential([]int{1, 3, 5, 7, 9}, 7) returns 3.
// Time complexity: O(log n). Additional space: O(1).
func Exponential[T cmp.Ordered](values []T, target T) int {
	return ExponentialFunc(values, target, orderedCompare[T])
}

// ExponentialFunc searches a sorted slice with a custom comparator using exponential search.
//
// The input must already be sorted according to compare. For duplicate values, ExponentialFunc
// returns the leftmost matching index.
// Example: ExponentialFunc(records, target, recordCompare).
// Time complexity: O(log n). Additional space: O(1).
func ExponentialFunc[T any](values []T, target T, compare func(a, b T) int) int {
	compare = requireCompare(compare)
	if len(values) == 0 {
		return -1
	}
	if compare(values[0], target) == 0 {
		return 0
	}

	bound := 1
	for bound < len(values) && compare(values[bound], target) < 0 {
		bound <<= 1
	}
	low := bound >> 1
	high := min(bound+1, len(values))
	index := lowerBound(values, target, compare, low, high)
	if index < len(values) && compare(values[index], target) == 0 {
		return index
	}
	return -1
}
