package sorting

import "cmp"

// Selection returns a sorted copy of values using selection sort.
//
// Selection sort performs a fixed number of comparisons regardless of input order.
// It is simple and predictable, but usually slower than insertion sort on small arrays.
// Example: Selection([]int{3, 1, 2}) returns []int{1, 2, 3}.
// Time complexity: O(n^2) in best, average, and worst case. Additional space: O(n).
func Selection[T cmp.Ordered](values []T) []T {
	return SelectionFunc(values, orderedLess[T])
}

// SelectionFunc returns a sorted copy of values using selection sort and the provided comparator.
//
// Use it when you want the selection-sort behavior for custom types.
// Example: SelectionFunc(items, func(a, b item) bool { return a.Priority < b.Priority }).
// Time complexity: O(n^2) in best, average, and worst case. Additional space: O(n).
func SelectionFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	SelectionInPlaceFunc(out, requireLess(less))
	return out
}

// SelectionInPlace sorts values in ascending order using selection sort.
//
// Selection sort minimizes swaps, which can be useful when writes are expensive, but it is not stable.
// Example: SelectionInPlace([]int{9, 4, 6}) changes the slice to []int{4, 6, 9}.
// Time complexity: O(n^2) in best, average, and worst case. Additional space: O(1).
func SelectionInPlace[T cmp.Ordered](values []T) {
	SelectionInPlaceFunc(values, orderedLess[T])
}

// SelectionInPlaceFunc sorts values in place using selection sort and the provided comparator.
//
// Example: SelectionInPlaceFunc(tasks, func(a, b task) bool { return a.Cost < b.Cost }).
// Time complexity: O(n^2) in best, average, and worst case. Additional space: O(1).
func SelectionInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	selectionInPlace(values, requireLess(less))
}

// selectionInPlace scans the unsorted suffix for its minimum element and moves it into place.
func selectionInPlace[T any](values []T, less func(a, b T) bool) {
	for i := 0; i < len(values)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(values); j++ {
			if less(values[j], values[minIndex]) {
				minIndex = j
			}
		}
		if minIndex != i {
			values[i], values[minIndex] = values[minIndex], values[i]
		}
	}
}
