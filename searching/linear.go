package searching

// Linear scans values from left to right and returns the index of the first matching target.
//
// Linear search is useful for unsorted data or very small slices where setup costs matter more
// than asymptotic performance.
// Example: Linear([]int{7, 3, 9, 3}, 3) returns 1.
// Time complexity: O(n) best, average, and worst case. Additional space: O(1).
func Linear[T comparable](values []T, target T) int {
	for index, value := range values {
		if value == target {
			return index
		}
	}
	return -1
}

// LinearFunc scans values from left to right and returns the index of the first matching target
// according to compare.
//
// Use it for custom types when equality should be defined by a comparator.
// Example: LinearFunc(users, target, func(a, b user) int { return cmp.Compare(a.ID, b.ID) }).
// Time complexity: O(n) best, average, and worst case. Additional space: O(1).
func LinearFunc[T any](values []T, target T, compare func(a, b T) int) int {
	compare = requireCompare(compare)
	for index, value := range values {
		if compare(value, target) == 0 {
			return index
		}
	}
	return -1
}
