package sorting

import "math/bits"

// Radix returns a sorted copy of integer values using radix sort.
//
// Radix sort is a non-comparison sort that can outperform comparison sorts on large integer-heavy
// inputs. This implementation handles signed ints by sorting a bit-transformed key that preserves
// normal integer ordering.
// Example: Radix([]int{170, -45, 75, -90}) returns []int{-90, -45, 75, 170}.
// Time complexity: O(w*n), where w is the number of processed bytes. Additional space: O(n).
func Radix(values []int) []int {
	out := clone(values)
	RadixInPlace(out)
	return out
}

// RadixInPlace sorts integer values in ascending order using radix sort.
//
// The algorithm performs a stable counting pass for each byte of the transformed integer key.
// Example: RadixInPlace([]int{12, -5, 7, 0}) changes the slice to []int{-5, 0, 7, 12}.
// Time complexity: O(w*n), where w is the number of processed bytes. Additional space: O(n).
func RadixInPlace(values []int) {
	if len(values) < 2 {
		return
	}

	buffer := make([]int, len(values))
	from := values
	to := buffer
	const radix = 256
	var counts [radix]int
	signMask := uint(1) << (bits.UintSize - 1)

	for shift := 0; shift < bits.UintSize; shift += 8 {
		for i := range counts {
			counts[i] = 0
		}
		for _, value := range from {
			key := uint(value) ^ signMask
			counts[(key>>shift)&0xFF]++
		}
		total := 0
		for i := range counts {
			total, counts[i] = total+counts[i], total
		}
		for _, value := range from {
			key := uint(value) ^ signMask
			bucket := (key >> shift) & 0xFF
			to[counts[bucket]] = value
			counts[bucket]++
		}
		from, to = to, from
	}

	if &from[0] != &values[0] {
		copy(values, from)
	}
}
