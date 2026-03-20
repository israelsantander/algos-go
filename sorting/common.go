package sorting

import "cmp"

const quickInsertionCutoff = 16

// clone returns a shallow copy of values while preserving nil slices as nil.
// It runs in O(n) time and uses O(n) additional space.
func clone[T any](values []T) []T {
	return append([]T(nil), values...)
}

// orderedLess is the default ascending comparator for ordered types.
// It is used to connect the ordered APIs to the generic Func variants.
func orderedLess[T cmp.Ordered](a, b T) bool {
	return a < b
}

// requireLess rejects nil comparators at the API boundary so failures are immediate and obvious.
func requireLess[T any](less func(a, b T) bool) func(a, b T) bool {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	return less
}
