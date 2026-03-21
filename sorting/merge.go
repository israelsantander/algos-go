package sorting

import (
	"cmp"
	"slices"
)

// Merge returns a sorted copy of values using merge sort.
//
// # Characteristics
//
// Merge sort is stable and guarantees O(n log n) time, making it a strong default when
// predictable performance matters more than temporary memory usage.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(n) additional space.
func Merge[T cmp.Ordered](values []T) []T {
	return MergeFunc(values, cmp.Less[T])
}

// MergeFunc returns a sorted copy of values using merge sort and the provided comparator.
//
// # Use
//
// Use it for custom types when you want a stable O(n log n) sort.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(n) additional space.
func MergeFunc[T any](values []T, less func(a, b T) bool) []T {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	out := slices.Clone(values)
	mergeInPlace(out, less)
	return out
}

// MergeInPlace sorts values in place using merge sort.
//
// # Characteristics
//
// The algorithm is stable but still needs an auxiliary buffer under the hood.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(n) additional space.
func MergeInPlace[T cmp.Ordered](values []T) {
	MergeInPlaceFunc(values, cmp.Less[T])
}

// MergeInPlaceFunc sorts values in place using merge sort and the provided comparator.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// It runs in O(n log n) time and uses O(n) additional space.
func MergeInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	mergeInPlace(values, less)
}

// mergeInPlace allocates one reusable buffer and starts the alternating source/destination recursion.
// This keeps the implementation stable without allocating a fresh slice at every recursion level.
func mergeInPlace[T any](values []T, less func(a, b T) bool) {
	n := len(values)
	if n < 2 {
		return
	}

	buffer := make([]T, n)
	copy(buffer, values)
	mergeSort(values, buffer, 0, n, less)
}

// mergeSort sorts src[start:end] into dst[start:end], swapping source and destination on recursive
// calls so one buffer can be reused throughout the entire sort.
func mergeSort[T any](dst, src []T, start, end int, less func(a, b T) bool) {
	if end-start <= 1 {
		if end-start == 1 {
			dst[start] = src[start]
		}
		return
	}

	if end-start <= quickInsertionCutoff {
		copy(dst[start:end], src[start:end])
		insertionInPlace(dst[start:end], less)
		return
	}

	mid := start + (end-start)/2
	mergeSort(src, dst, start, mid, less)
	mergeSort(src, dst, mid, end, less)

	// When the two runs are already in order, the merge step collapses to one copy.
	if !less(src[mid], src[mid-1]) {
		copy(dst[start:end], src[start:end])
		return
	}

	mergeRuns(dst, src, start, mid, end, less)
}

// mergeRuns performs the linear merge step that gives merge sort its stable behavior.
func mergeRuns[T any](dst, src []T, start, mid, end int, less func(a, b T) bool) {
	left, right := start, mid
	for write := start; write < end; write++ {
		switch {
		case left >= mid:
			dst[write] = src[right]
			right++
		case right >= end:
			dst[write] = src[left]
			left++
		case less(src[right], src[left]):
			dst[write] = src[right]
			right++
		default:
			dst[write] = src[left]
			left++
		}
	}
}
