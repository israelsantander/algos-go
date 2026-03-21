package sorting

import "slices"

// Counting returns a sorted copy of integer values using counting sort.
//
// # Characteristics
//
// Counting sort is fastest when the input values span a relatively compact integer range.
//
// This implementation supports negative integers by offsetting counts by the observed minimum.
//
// # Complexity
//
// It runs in O(n + k) time, where k is the value range, and uses O(n + k) additional space.
func Counting(values []int) []int {
	out := slices.Clone(values)
	CountingInPlace(out)
	return out
}

// CountingInPlace sorts integer values in place using counting sort.
//
// # Characteristics
//
// Counting sort is stable in concept, but this API rewrites the input slice from accumulated
// counts rather than preserving the original order of equal values.
//
// # Complexity
//
// It runs in O(n + k) time, where k is the value range, and uses O(k) additional space.
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
