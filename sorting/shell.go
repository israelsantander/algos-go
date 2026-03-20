package sorting

import "cmp"

// Shell returns a sorted copy of values using Shell sort.
//
// Shell sort generalizes insertion sort by first moving elements across larger gaps, which makes it
// much faster than plain quadratic sorts on medium-sized slices.
// Example: Shell([]int{8, 5, 2, 6, 9, 3}) returns []int{2, 3, 5, 6, 8, 9}.
// Time complexity: depends on the gap sequence; with Knuth gaps it is typically subquadratic in practice.
// Additional space: O(n).
func Shell[T cmp.Ordered](values []T) []T {
	return ShellFunc(values, orderedLess[T])
}

// ShellFunc returns a sorted copy of values using Shell sort and the provided comparator.
//
// Use it for custom types when you want an in-place algorithm that often beats insertion and selection sort.
// Example: ShellFunc(items, func(a, b item) bool { return a.Weight < b.Weight }).
// Time complexity: depends on the gap sequence; with Knuth gaps it is typically subquadratic in practice.
// Additional space: O(n).
func ShellFunc[T any](values []T, less func(a, b T) bool) []T {
	out := clone(values)
	ShellInPlaceFunc(out, requireLess(less))
	return out
}

// ShellInPlace sorts values in ascending order using Shell sort.
//
// It is in-place and usually much faster than simple quadratic sorts, but it is not stable.
// Example: ShellInPlace([]int{7, 1, 4, 2}) changes the slice to []int{1, 2, 4, 7}.
// Time complexity: depends on the gap sequence; with Knuth gaps it is typically subquadratic in practice.
// Additional space: O(1).
func ShellInPlace[T cmp.Ordered](values []T) {
	ShellInPlaceFunc(values, orderedLess[T])
}

// ShellInPlaceFunc sorts values in place using Shell sort and the provided comparator.
//
// Example: ShellInPlaceFunc(entries, func(a, b entry) bool { return a.Key < b.Key }).
// Time complexity: depends on the gap sequence; with Knuth gaps it is typically subquadratic in practice.
// Additional space: O(1).
func ShellInPlaceFunc[T any](values []T, less func(a, b T) bool) {
	shellInPlace(values, requireLess(less))
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
