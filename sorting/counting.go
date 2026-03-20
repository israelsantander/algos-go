package sorting

// Counting returns a sorted copy of integer values using counting sort.
//
// Counting sort is fastest when the input values span a relatively compact integer range.
// This implementation supports negative integers by offsetting counts by the observed minimum.
// Example: Counting([]int{4, -1, 2, -1}) returns []int{-1, -1, 2, 4}.
// Time complexity: O(n + k), where k is the value range. Additional space: O(n + k).
func Counting(values []int) []int {
	out := clone(values)
	CountingInPlace(out)
	return out
}

// CountingInPlace sorts integer values in ascending order using counting sort.
//
// Counting sort is stable in concept, but this in-place-style API rewrites the input slice from
// accumulated counts rather than preserving original equal-value order.
// Example: CountingInPlace([]int{3, -2, 3, 1}) changes the slice to []int{-2, 1, 3, 3}.
// Time complexity: O(n + k), where k is the value range. Additional space: O(k).
func CountingInPlace(values []int) {
	if len(values) < 2 {
		return
	}

	minValue, maxValue := values[0], values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}

	rangeSize := maxValue - minValue + 1
	counts := make([]int, rangeSize)
	for _, value := range values {
		counts[value-minValue]++
	}

	write := 0
	for offset, count := range counts {
		value := minValue + offset
		for count > 0 {
			values[write] = value
			write++
			count--
		}
	}
}
