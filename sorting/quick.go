package sorting

import (
	"cmp"
	"slices"
)

const quickInsertionCutoff = 16

// Quick returns a sorted copy of values using quicksort.
//
// # Characteristics
//
// Quicksort is often one of the fastest general-purpose in-memory sorts.
//
// This implementation uses median-of-three pivoting and 2-way partitioning (Hoare's scheme)
// with an insertion-sort cutoff for small partitions.
//
// It is not stable.
//
// # Complexity
//
// It runs in O(n log n) time on average, O(n^2) in the worst case, and uses O(n) additional
// space for the copied result plus O(log n) stack space.
func Quick[T cmp.Ordered](values []T) []T {
	return QuickFunc(values, cmp.Less[T])
}

// QuickFunc returns a sorted copy of values using quicksort and the provided comparator.
//
// # Use
//
// Use it for custom types when you want a very fast general-purpose comparison sort.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n log n) time on average, O(n^2) in the worst case, and uses O(n) additional
// space for the copied result plus O(log n) stack space.
func QuickFunc[T any](values []T, less func(a, b T) bool) []T {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	out := slices.Clone(values)
	quickInPlace(out, less)
	return out
}

// QuickInPlace sorts values in ascending order using quicksort.
//
// # Characteristics
//
// It is usually faster than the simpler quadratic algorithms, but it is not stable.
//
// # Complexity
//
// It runs in O(n log n) time on average, O(n^2) in the worst case, and uses O(log n) stack space.
func QuickInPlace[T cmp.Ordered](values []T) {
	QuickInPlaceFunc(values, cmp.Less[T])
}

// QuickInPlaceFunc sorts values in place using quicksort and the provided comparator.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n log n) time on average, O(n^2) in the worst case, and uses O(log n) stack space.
func QuickInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	quickInPlace(values, less)
}

// quickInPlace handles the trivial cases before delegating to the partitioning helper.
func quickInPlace[T any](values []T, less func(a, b T) bool) {
	if len(values) < 2 {
		return
	}

	quickSort(values, 0, len(values)-1, less)
}

// quickSort uses insertion sort on tiny partitions, median-of-three pivot selection, and Hoare's
// partitioning scheme. It recurses into the smaller side first and loops over the larger side to bound stack use.
func quickSort[T any](values []T, low, high int, less func(a, b T) bool) {
	for low < high {
		if high-low+1 <= quickInsertionCutoff {
			insertionInPlace(values[low:high+1], less)
			return
		}

		pivotIndex := medianOfThree(values, low, high, less)
		p := partitionHoare(values, low, high, pivotIndex, less)

		if p-low < high-p {
			quickSort(values, low, p, less)
			low = p + 1
		} else {
			quickSort(values, p+1, high, less)
			high = p
		}
	}
}

// medianOfThree chooses a pivot index from the low, middle, and high elements to reduce the chance
// of consistently poor partitions on already-structured input. It also sorts these three
// elements in place to slightly improve the partitioning step.
func medianOfThree[T any](values []T, low, high int, less func(a, b T) bool) int {
	mid := low + (high-low)/2

	if less(values[mid], values[low]) {
		values[low], values[mid] = values[mid], values[low]
	}
	if less(values[high], values[low]) {
		values[low], values[high] = values[high], values[low]
	}
	if less(values[high], values[mid]) {
		values[mid], values[high] = values[high], values[mid]
	}

	return mid
}

// medianOfThreeValue is like medianOfThree but returns the pivot value itself.
func medianOfThreeValue[T any](values []T, low, high int, less func(a, b T) bool) T {
	mid := medianOfThree(values, low, high, less)
	return values[mid]
}

// partitionHoare uses Hoare's scheme, which is generally more efficient than Lomuto's (as it
// performs fewer swaps) and handles already-sorted arrays well.
func partitionHoare[T any](values []T, low, high int, pivotIndex int, less func(a, b T) bool) int {
	pivot := values[pivotIndex]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if !less(values[i], pivot) {
				break
			}
		}
		for {
			j--
			if !less(pivot, values[j]) {
				break
			}
		}
		if i >= j {
			return j
		}
		values[i], values[j] = values[j], values[i]
	}
}

// partition3Way groups values into less-than, equal-to, and greater-than pivot regions in one pass.
// That keeps duplicates from degrading performance as badly as a simple 2-way partition would.
func partition3Way[T any](values []T, low, high int, pivot T, less func(a, b T) bool) (int, int) {
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
