package sorting

import "cmp"

// Heap returns a sorted copy of values using heap sort.
//
// Heap sort guarantees O(n log n) time and uses only constant extra working memory in its in-place form.
// It is a good fit when worst-case time matters and stability does not.
// Example: Heap([]int{6, 1, 8, 3}) returns []int{1, 3, 6, 8}.
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(n).
func Heap[T cmp.Ordered](values []T) []T {
	return HeapFunc(values, orderedLess[T])
}

// HeapFunc returns a sorted copy of values using heap sort and the provided comparator.
//
// Use it for custom types when you want deterministic O(n log n) sorting without merge sort's extra buffer logic.
// Example: HeapFunc(files, func(a, b file) bool { return a.Size < b.Size }).
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(n).
func HeapFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	HeapInPlaceFunc(out, requireLess(less))
	return out
}

// HeapInPlace sorts values in ascending order using heap sort.
//
// It is in-place and guarantees O(n log n) time, but it is not stable and often has larger
// constant factors than quicksort in practice.
// Example: HeapInPlace([]int{5, 9, 1, 3}) changes the slice to []int{1, 3, 5, 9}.
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(1).
func HeapInPlace[T cmp.Ordered](values []T) {
	HeapInPlaceFunc(values, orderedLess[T])
}

// HeapInPlaceFunc sorts values in place using heap sort and the provided comparator.
//
// Example: HeapInPlaceFunc(items, func(a, b item) bool { return a.Rank < b.Rank }).
// Time complexity: O(n log n) in best, average, and worst case. Additional space: O(1).
func HeapInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	heapInPlace(values, requireLess(less))
}

// heapInPlace first builds a max-heap, then repeatedly moves the largest element to the end
// of the slice and shrinks the heap boundary.
func heapInPlace[T any](values []T, less func(a, b T) bool) {
	n := len(values)
	for root := n/2 - 1; root >= 0; root-- {
		siftDown(values, root, n, less)
	}

	for end := n - 1; end > 0; end-- {
		values[0], values[end] = values[end], values[0]
		siftDown(values, 0, end, less)
	}
}

// siftDown restores the max-heap property by pushing one out-of-place root downward until
// it dominates both of its children again.
func siftDown[T any](values []T, root, size int, less func(a, b T) bool) {
	for {
		left := 2*root + 1
		if left >= size {
			return
		}

		largest := left
		right := left + 1
		if right < size && less(values[largest], values[right]) {
			largest = right
		}
		if !less(values[root], values[largest]) {
			return
		}

		values[root], values[largest] = values[largest], values[root]
		root = largest
	}
}
