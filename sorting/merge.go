package sorting

import "cmp"

// Merge returns a sorted copy of values using merge sort.
//
// Merge sort is stable and guarantees O(n log n) time, making it a strong default when predictable
// performance matters more than temporary memory usage.
// Example: Merge([]int{5, 2, 5, 1}) returns []int{1, 2, 5, 5}.
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(n).
func Merge[T cmp.Ordered](values []T) []T {
	return MergeFunc(values, orderedLess[T])
}

// MergeFunc returns a sorted copy of values using merge sort and the provided comparator.
//
// Use it for custom types when you want a stable O(n log n) sort.
// Example: MergeFunc(users, func(a, b user) bool { return a.Name < b.Name }).
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(n).
func MergeFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	MergeInPlaceFunc(out, requireLess(less))
	return out
}

// MergeInPlace sorts values in ascending order using merge sort.
//
// The input slice is overwritten with the sorted result. The algorithm is stable but still needs
// an auxiliary buffer under the hood.
// Example: MergeInPlace([]int{4, 1, 4, 2}) changes the slice to []int{1, 2, 4, 4}.
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(n).
func MergeInPlace[T cmp.Ordered](values []T) {
	MergeInPlaceFunc(values, orderedLess[T])
}

// MergeInPlaceFunc sorts values in place using merge sort and the provided comparator.
//
// Example: MergeInPlaceFunc(records, func(a, b record) bool { return a.Timestamp < b.Timestamp }).
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(n).
func MergeInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	mergeInPlace(values, requireLess(less))
}

// mergeInPlace allocates one reusable buffer and starts the alternating source/destination recursion.
// This keeps the implementation stable without allocating a fresh slice at every recursion level.
func mergeInPlace[T any](values []T, less func(a, b T) bool) {
	if len(values) < 2 {
		return
	}

	buffer := clone(values)
	mergeSort(values, buffer, 0, len(values), less)
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
