package sorting

// Reverse returns a reversed copy of values.
//
// Use it when you want the input left untouched but need the elements in the opposite order.
// Example: Reverse([]int{1, 2, 3}) returns []int{3, 2, 1}.
// Time complexity: O(n). Additional space: O(n).
func Reverse[T any](values []T) []T {
	out := clone(values)
	reverseInPlace(out)
	return out
}

// ReverseInPlace reverses values in place.
//
// Use it when mutation is acceptable and you want the smallest possible memory footprint.
// Example: ReverseInPlace([]int{1, 2, 3}) changes the slice to []int{3, 2, 1}.
// Time complexity: O(n). Additional space: O(1).
func ReverseInPlace[T any](values []T) {
	reverseInPlace(values)
}

// reverseInPlace swaps mirrored elements until both pointers meet in the middle.
func reverseInPlace[T any](values []T) {
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		values[left], values[right] = values[right], values[left]
	}
}
