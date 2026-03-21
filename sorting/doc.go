// Package sorting provides educational implementations of classic sorting algorithms.
//
// # API Families
//
// Ordered variants such as Bubble and QuickInPlace sort cmp.Ordered values in ascending order.
// Integer-specific variants such as Counting and Radix focus on dense numeric workloads.
// Func variants accept a custom comparator for arbitrary element types or custom orderings.
//
// # Comparator Requirements
//
// For all Func variants, less must report whether a should sort before b.
//
// It must define a strict weak ordering and must not mutate its arguments or the slice being
// sorted. Passing a nil comparator panics. Results are unspecified if less is inconsistent or
// has side effects.
package sorting
