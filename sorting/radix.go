package sorting

import "math/bits"
import "slices"

const radix uint = 256
// Radix returns a sorted copy of integer values using radix sort.
//
// # Characteristics
//
// Radix sort is a non-comparison sort that can outperform comparison sorts on large integer-heavy
// inputs. This implementation handles signed ints by sorting a bit-transformed key that preserves
// normal integer ordering.
//
// # Complexity
//
// It runs in O(w*n) time, where w is the number of processed bytes, and uses O(n) additional space.
func Radix(values []int) []int {
	out := slices.Clone(values)
	RadixInPlace(out)
	return out
}

// RadixInPlace sorts integer values in place using radix sort.
//
// # Characteristics
//
// The algorithm performs a stable counting pass for each byte of the transformed integer key.
//
// # Complexity
//
// It runs in O(w*n) time, where w is the number of processed bytes, and uses O(n) additional space.
func RadixInPlace(values []int) {
	if len(values) < 2 {
		return
	}

	buffer := make([]int, len(values))
	from := values
	to := buffer
	
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
