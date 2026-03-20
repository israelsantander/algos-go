package searching

import "cmp"

// clone returns a shallow copy of values while preserving nil slices as nil.
func clone[T any](values []T) []T {
	return append([]T(nil), values...)
}

// orderedCompare is the default ascending comparator for ordered types.
func orderedCompare[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

// requireCompare rejects nil comparators at the API boundary so failures are immediate and obvious.
func requireCompare[T any](compare func(a, b T) int) func(a, b T) int {
	if compare == nil {
		panic("searching: compare function is nil")
	}
	return compare
}
