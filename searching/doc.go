// Package searching provides educational implementations of classic search algorithms.
//
// Ordered variants such as Binary, Exponential, Jump, SearchBST, and AVLSearch work with
// cmp.Ordered values.
// Func variants accept a tri-state comparator for custom element types or custom orderings.
//
// For all Func variants, compare must return:
//   - a negative value when a should sort before b
//   - zero when a and b compare equal
//   - a positive value when a should sort after b
//
// Passing a nil comparator panics. Results are unspecified if compare is inconsistent or
// mutates the values being searched.
package searching
