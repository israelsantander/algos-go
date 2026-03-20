package searching

import (
	"cmp"
	"slices"
)

// Binary searches a sorted slice and returns the index of the first matching target.
//
// The input must already be sorted in ascending order. For duplicate values, Binary returns the
// leftmost matching index.
// Example: Binary([]int{1, 3, 3, 5}, 3) returns 1.
// Time complexity: O(log n). Additional space: O(1).
func Binary[T cmp.Ordered](values []T, target T) int {
	return BinaryFunc(values, target, orderedCompare[T])
}

// BinaryFunc searches a sorted slice and returns the index of the first matching target according
// to compare.
//
// The input must already be sorted according to compare. For duplicate values, BinaryFunc returns
// the leftmost matching index.
// Example: BinaryFunc(users, target, userCmp).
// Time complexity: O(log n). Additional space: O(1).
func BinaryFunc[T any](values []T, target T, compare func(a, b T) int) int {
	compare = requireCompare(compare)
	index := lowerBound(values, target, compare, 0, len(values))
	if index < len(values) && compare(values[index], target) == 0 {
		return index
	}
	return -1
}

// SortedForBinary returns a sorted copy of values for use with Binary.
//
// Use it when you want to keep the original slice unchanged before running repeated binary searches.
// Example: SortedForBinary([]int{9, 1, 4}) returns []int{1, 4, 9}.
// Time complexity: O(n log n). Additional space: O(n).
func SortedForBinary[T cmp.Ordered](values []T) []T {
	return SortedForBinaryFunc(values, orderedCompare[T])
}

// SortedForBinaryFunc returns a sorted copy of values for use with BinaryFunc.
//
// Use it for custom types or custom orderings where binary search requires a comparator.
// Example: SortedForBinaryFunc(users, userCmp).
// Time complexity: O(n log n). Additional space: O(n).
func SortedForBinaryFunc[T any](values []T, compare func(a, b T) int) []T {
	out := clone(values)
	slices.SortFunc(out, requireCompare(compare))
	return out
}

// lowerBound returns the first index in values[low:high] whose value is not less than target.
func lowerBound[T any](values []T, target T, compare func(a, b T) int, low, high int) int {
	for low < high {
		mid := low + (high-low)/2
		if compare(values[mid], target) < 0 {
			low = mid + 1
			continue
		}
		high = mid
	}
	return low
}
