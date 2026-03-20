package sorting

import "cmp"

// Bubble returns a sorted copy of values using bubble sort.
//
// Bubble sort is mainly useful for teaching or very small inputs. This version stops early
// when the slice is already sorted or nearly sorted.
// Example: Bubble([]int{5, 1, 4, 2}) returns []int{1, 2, 4, 5}.
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(n).
func Bubble[T cmp.Ordered](values []T) []T {
	return BubbleFunc(values, orderedLess[T])
}

// BubbleFunc returns a sorted copy of values using bubble sort and the provided comparator.
//
// Use it for custom element types or custom orderings.
// Example: BubbleFunc(users, func(a, b user) bool { return a.Age < b.Age }).
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(n).
func BubbleFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	BubbleInPlaceFunc(out, requireLess(less))
	return out
}

// BubbleInPlace sorts values in ascending order using bubble sort.
//
// It is stable and in-place, but still quadratic, so it is best reserved for small or educational cases.
// Example: BubbleInPlace([]int{3, 2, 1}) changes the slice to []int{1, 2, 3}.
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(1).
func BubbleInPlace[T cmp.Ordered](values []T) {
	BubbleInPlaceFunc(values, orderedLess[T])
}

// BubbleInPlaceFunc sorts values in place using bubble sort and the provided comparator.
//
// Example: BubbleInPlaceFunc(records, func(a, b record) bool { return a.Score < b.Score }).
// Time complexity: O(n^2) average and worst case, O(n) best case. Additional space: O(1).
func BubbleInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	bubbleInPlace(values, requireLess(less))
}

// bubbleInPlace tracks the last swap position so each pass only revisits the unsorted suffix.
// That early-exit optimization makes already-sorted input finish in linear time.
func bubbleInPlace[T any](values []T, less func(a, b T) bool) {
	for n := len(values); n > 1; {
		lastSwap := 0
		for i := 1; i < n; i++ {
			if less(values[i], values[i-1]) {
				values[i-1], values[i] = values[i], values[i-1]
				lastSwap = i
			}
		}
		if lastSwap == 0 {
			return
		}
		n = lastSwap
	}
}
