package sorting

import (
	"cmp"
	"slices"
)

// Insertion returns a sorted copy of values using insertion sort.
//
// # Characteristics
//
// Insertion sort is a strong choice for tiny inputs and nearly sorted data because it moves each
// element only as far as needed. It is stable.
//
// # Complexity
//
// It runs in O(n^2) time in the average and worst cases, O(n) on already-sorted input,
// and uses O(n) additional space for the copied result.
func Insertion[T cmp.Ordered](values []T) []T {
	return InsertionFunc(values, cmp.Less[T])
}

// InsertionFunc returns a sorted copy of values using insertion sort and the provided comparator.
//
// # Use
//
// Use it for custom types when the data is small or already close to sorted.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n^2) time in the average and worst cases, O(n) on already-sorted input,
// and uses O(n) additional space for the copied result.
func InsertionFunc[T any](values []T, less func(a, b T) bool) []T {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	out := slices.Clone(values)
	insertionInPlace(out, less)
	return out
}

// InsertionInPlace sorts values in place using insertion sort.
//
// # Characteristics
//
// Insertion sort is stable, in-place, and especially effective on nearly sorted slices.
//
// # Complexity
//
// It runs in O(n^2) time in the average and worst cases, O(n) on already-sorted input,
// and uses O(1) additional space.
func InsertionInPlace[T cmp.Ordered](values []T) {
	InsertionInPlaceFunc(values, cmp.Less[T])
}

// InsertionInPlaceFunc sorts values in place using insertion sort and the provided comparator.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n^2) time in the average and worst cases, O(n) on already-sorted input,
// and uses O(1) additional space.
func InsertionInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	insertionInPlace(values, less)
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
