package searching

import (
	"cmp"
	"math"
)

// Jump searches a sorted slice using fixed-size blocks followed by a short linear scan.
//
// Jump search trades a little extra work against simpler memory access patterns than binary search.
// For duplicate values, Jump returns the leftmost matching index.
// Example: Jump([]int{1, 3, 5, 7, 9}, 7) returns 3.
// Time complexity: O(sqrt(n)). Additional space: O(1).
func Jump[T cmp.Ordered](values []T, target T) int {
	return JumpFunc(values, target, orderedCompare[T])
}

// JumpFunc searches a sorted slice using jump search and a custom comparator.
//
// The input must already be sorted according to compare. For duplicate values, JumpFunc returns
// the leftmost matching index.
// Example: JumpFunc(records, target, recordCompare).
// Time complexity: O(sqrt(n)). Additional space: O(1).
func JumpFunc[T any](values []T, target T, compare func(a, b T) int) int {
	compare = requireCompare(compare)
	if len(values) == 0 {
		return -1
	}

	step := int(math.Sqrt(float64(len(values))))
	if step < 1 {
		step = 1
	}
	low := 0
	high := min(step, len(values))
	for low < len(values) && compare(values[high-1], target) < 0 {
		low = high
		if low >= len(values) {
			return -1
		}
		high = min(high+step, len(values))
	}

	index := lowerBound(values, target, compare, low, high)
	if index < len(values) && compare(values[index], target) == 0 {
		return index
	}
	return -1
}
