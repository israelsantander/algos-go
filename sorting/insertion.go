package sorting

import "cmp"

// Insertion returns a sorted copy of values using insertion sort.
//
// Insertion sort is a strong choice for tiny inputs and nearly sorted data because it moves each
// element only as far as needed.
// Example: Insertion([]int{4, 1, 3, 2}) returns []int{1, 2, 3, 4}.
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(n).
func Insertion[T cmp.Ordered](values []T) []T {
	return InsertionFunc(values, orderedLess[T])
}

// InsertionFunc returns a sorted copy of values using insertion sort and the provided comparator.
//
// Use it for custom types when the data is small or already close to sorted.
// Example: InsertionFunc(points, func(a, b point) bool { return a.X < b.X }).
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(n).
func InsertionFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	InsertionInPlaceFunc(out, requireLess(less))
	return out
}

// InsertionInPlace sorts values in ascending order using insertion sort.
//
// It is stable, in-place, and especially effective on nearly sorted slices.
// Example: InsertionInPlace([]int{2, 1, 3}) changes the slice to []int{1, 2, 3}.
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(1).
func InsertionInPlace[T cmp.Ordered](values []T) {
	InsertionInPlaceFunc(values, orderedLess[T])
}

// InsertionInPlaceFunc sorts values in place using insertion sort and the provided comparator.
//
// Example: InsertionInPlaceFunc(rows, func(a, b row) bool { return a.ID < b.ID }).
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(1).
func InsertionInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	insertionInPlace(values, requireLess(less))
}

// insertionInPlace keeps a sorted prefix and inserts the next value into its correct spot.
func insertionInPlace[T any](values []T, less func(a, b T) bool) {
	for i := 1; i < len(values); i++ {
		current := values[i]
		j := i
		for ; j > 0 && less(current, values[j-1]); j-- {
			values[j] = values[j-1]
		}
		values[j] = current
	}
}
