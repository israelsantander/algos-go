package sorting

import "slices"

// Reverse returns a reversed copy of values.
//
// # Characteristics
//
// Use it when you want the input left untouched but need the elements in the opposite order.
//
// # Complexity
//
// It runs in O(n) time and uses O(n) additional space.
func Reverse[T any](values []T) []T {
	out := slices.Clone(values)
	reverseInPlace(out)
	return out
}

// ReverseInPlace reverses values in place.
//
// # Characteristics
//
// Use it when mutation is acceptable and you want the smallest possible memory footprint.
//
// # Complexity
//
// It runs in O(n) time and uses O(1) additional space.
func ReverseInPlace[T any](values []T) {
	reverseInPlace(values)
}

// reverseInPlace swaps mirrored elements until both pointers meet in the middle.
func reverseInPlace[T any](values []T) {
	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		values[left], values[right] = values[right], values[left]
	}
}
