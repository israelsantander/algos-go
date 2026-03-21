package sorting

import (
	"cmp"
	"slices"
)

// Shell returns a sorted copy of values using Shell sort.
//
// # Characteristics
//
// Shell sort generalizes insertion sort by first moving elements across larger gaps, which makes
// it much faster than plain quadratic sorts on medium-sized slices. It is not stable.
//
// # Complexity
//
// Its exact running time depends on the gap sequence.
//
// With Knuth gaps, it is typically subquadratic in practice and uses O(n) additional space.
func Shell[T cmp.Ordered](values []T) []T {
	return ShellFunc(values, cmp.Less[T])
}

// ShellFunc returns a sorted copy of values using Shell sort and the provided comparator.
//
// # Use
//
// Use it for custom types when you want an in-place algorithm that often beats insertion and selection sort.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// Its exact running time depends on the gap sequence.
//
// With Knuth gaps, it is typically subquadratic in practice and uses O(n) additional space.
func ShellFunc[T any](values []T, less func(a, b T) bool) []T {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	out := slices.Clone(values)
	shellInPlace(out, less)
	return out
}

// ShellInPlace sorts values in place using Shell sort.
//
// # Characteristics
//
// Shell sort is in-place and usually much faster than simple quadratic sorts, but it is not stable.
//
// # Complexity
//
// Its exact running time depends on the gap sequence.
//
// With Knuth gaps, it is typically subquadratic in practice and uses O(1) additional space.
func ShellInPlace[T cmp.Ordered](values []T) {
	ShellInPlaceFunc(values, cmp.Less[T])
}

// ShellInPlaceFunc sorts values in place using Shell sort and the provided comparator.
//
// # Requirements
//
// less must define a strict weak ordering.
//
// It panics if less is nil.
//
// # Complexity
//
// Its exact running time depends on the gap sequence.
//
// With Knuth gaps, it is typically subquadratic in practice and uses O(1) additional space.
func ShellInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	if less == nil {
		panic("sorting: less comparator is nil")
	}
	shellInPlace(values, less)
}

// shellInPlace performs gap-based insertion passes using the Knuth sequence, gradually reducing
// long-distance disorder before finishing with a standard gap-1 insertion pass.
func shellInPlace[T any](values []T, less func(a, b T) bool) {
	for gap := knuthGap(len(values)); gap > 0; gap /= 3 {
		for i := gap; i < len(values); i++ {
			current := values[i]
			j := i
			for ; j >= gap && less(current, values[j-gap]); j -= gap {
				values[j] = values[j-gap]
			}
			values[j] = current
		}
	}
}

// knuthGap returns the largest Knuth gap less than n so the outer loop can walk the sequence backward.
func knuthGap(n int) int {
	gap := 1
	for gap < n/3 {
		gap = 3*gap + 1
	}
	return gap
}
