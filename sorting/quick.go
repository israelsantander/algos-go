package sorting

import "cmp"

// Quick returns a sorted copy of values using quicksort.
//
// Quicksort is often one of the fastest general-purpose in-memory sorts. This implementation improves
// the classic algorithm with median-of-three pivoting, 3-way partitioning, and an insertion-sort cutoff.
// Example: Quick([]int{9, 3, 7, 1}) returns []int{1, 3, 7, 9}.
// Time complexity: O(n log n) average, O(n^2) worst case. Additional space: O(n) for the copied result plus O(log n) stack.
func Quick[T cmp.Ordered](values []T) []T {
	return QuickFunc(values, orderedLess[T])
}

// QuickFunc returns a sorted copy of values using quicksort and the provided comparator.
//
// Use it for custom types when you want a very fast general-purpose comparison sort.
// Example: QuickFunc(users, func(a, b user) bool { return a.Score < b.Score }).
// Time complexity: O(n log n) average, O(n^2) worst case. Additional space: O(n) for the copied result plus O(log n) stack.
func QuickFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	QuickInPlaceFunc(out, requireLess(less))
	return out
}

// QuickInPlace sorts values in ascending order using quicksort.
//
// It is usually faster than the simpler quadratic algorithms, but it is not stable.
// Example: QuickInPlace([]int{10, 4, 6, 2}) changes the slice to []int{2, 4, 6, 10}.
// Time complexity: O(n log n) average, O(n^2) worst case. Additional space: O(log n) stack.
func QuickInPlace[T cmp.Ordered](values []T) {
	QuickInPlaceFunc(values, orderedLess[T])
}

// QuickInPlaceFunc sorts values in place using quicksort and the provided comparator.
//
// Example: QuickInPlaceFunc(points, func(a, b point) bool { return a.Y < b.Y }).
// Time complexity: O(n log n) average, O(n^2) worst case. Additional space: O(log n) stack.
func QuickInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	quickInPlace(values, requireLess(less))
}

// quickInPlace handles the trivial cases before delegating to the partitioning helper.
func quickInPlace[T any](values []T, less func(a, b T) bool) {
	if len(values) < 2 {
		return
	}

	quickSort(values, 0, len(values)-1, less)
}

// quickSort uses insertion sort on tiny partitions, median-of-three pivot selection, and 3-way
// partitioning. It recurses into the smaller side first and loops over the larger side to bound stack use.
func quickSort[T any](values []T, low, high int, less func(a, b T) bool) {
	for low < high {
		if high-low+1 <= quickInsertionCutoff {
			insertionInPlace(values[low:high+1], less)
			return
		}

		pivot := medianOfThree(values, low, high, less)
		lt, gt := partition(values, low, high, pivot, less)

		if lt-low < high-gt {
			quickSort(values, low, lt-1, less)
			low = gt + 1
			continue
		}

		quickSort(values, gt+1, high, less)
		high = lt - 1
	}
}

// medianOfThree chooses a pivot from the low, middle, and high elements to reduce the chance
// of consistently poor partitions on already-structured input.
func medianOfThree[T any](values []T, low, high int, less func(a, b T) bool) T {
	mid := low + (high-low)/2

	if less(values[mid], values[low]) {
		values[low], values[mid] = values[mid], values[low]
	}
	if less(values[high], values[mid]) {
		values[mid], values[high] = values[high], values[mid]
	}
	if less(values[mid], values[low]) {
		values[low], values[mid] = values[mid], values[low]
	}

	return values[mid]
}

// partition groups values into less-than, equal-to, and greater-than pivot regions in one pass.
// That keeps duplicates from degrading performance as badly as a simple 2-way partition would.
func partition[T any](values []T, low, high int, pivot T, less func(a, b T) bool) (int, int) {
	lt, i, gt := low, low, high
	for i <= gt {
		switch {
		case less(values[i], pivot):
			values[lt], values[i] = values[i], values[lt]
			lt++
			i++
		case less(pivot, values[i]):
			values[i], values[gt] = values[gt], values[i]
			gt--
		default:
			i++
		}
	}
	return lt, gt
}
