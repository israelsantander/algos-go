package searching

import "cmp"

// BSTNode is a node in a binary search tree.
//
// The tree is intentionally unbalanced, so performance depends on insertion order.
type BSTNode[T any] struct {
	Value T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

// Insert adds value to the binary search tree rooted at root and returns the tree root.
//
// Duplicate values are inserted into the right subtree to keep behavior deterministic.
// Time complexity: O(h), where h is the tree height. Worst case: O(n).
func Insert[T cmp.Ordered](root *BSTNode[T], value T) *BSTNode[T] {
	return InsertFunc(root, value, orderedCompare[T])
}

// InsertFunc adds value to the binary search tree rooted at root using compare and returns the tree root.
//
// Duplicate values are inserted into the right subtree to keep behavior deterministic.
// Time complexity: O(h), where h is the tree height. Worst case: O(n).
func InsertFunc[T any](root *BSTNode[T], value T, compare func(a, b T) int) *BSTNode[T] {
	return insert(root, value, requireCompare(compare))
}

// insert adds value to the binary search tree rooted at root using a validated comparator.
func insert[T any](root *BSTNode[T], value T, compare func(a, b T) int) *BSTNode[T] {
	if root == nil {
		return &BSTNode[T]{Value: value}
	}

	current := root
	for {
		if compare(value, current.Value) < 0 {
			if current.Left == nil {
				current.Left = &BSTNode[T]{Value: value}
				return root
			}
			current = current.Left
			continue
		}
		if current.Right == nil {
			current.Right = &BSTNode[T]{Value: value}
			return root
		}
		current = current.Right
	}
}

// BuildBST inserts values in order and returns the resulting binary search tree root.
//
// Because the tree is unbalanced, sorted input produces the worst-case tree shape.
// Time complexity: O(n log n) average, O(n^2) worst case.
func BuildBST[T cmp.Ordered](values []T) *BSTNode[T] {
	return BuildBSTFunc(values, orderedCompare[T])
}

// BuildBSTFunc inserts values in order using compare and returns the resulting binary search tree root.
//
// Because the tree is unbalanced, sorted input produces the worst-case tree shape.
// Time complexity: O(n log n) average, O(n^2) worst case.
func BuildBSTFunc[T any](values []T, compare func(a, b T) int) *BSTNode[T] {
	compare = requireCompare(compare)
	var root *BSTNode[T]
	for _, value := range values {
		root = insert(root, value, compare)
	}
	return root
}

// SearchBST searches the binary search tree rooted at root and returns the matching node or nil.
//
// The tree must have been built with the same ordering used by Insert.
// Time complexity: O(h), where h is the tree height. Worst case: O(n).
func SearchBST[T cmp.Ordered](root *BSTNode[T], target T) *BSTNode[T] {
	return SearchBSTFunc(root, target, orderedCompare[T])
}

// SearchBSTFunc searches the binary search tree rooted at root using compare and returns the
// matching node or nil.
//
// The tree must have been built with the same ordering used by InsertFunc.
// Time complexity: O(h), where h is the tree height. Worst case: O(n).
func SearchBSTFunc[T any](root *BSTNode[T], target T, compare func(a, b T) int) *BSTNode[T] {
	return searchBST(root, target, requireCompare(compare))
}

// searchBST traverses the binary search tree with a validated comparator.
func searchBST[T any](root *BSTNode[T], target T, compare func(a, b T) int) *BSTNode[T] {
	for root != nil {
		switch result := compare(target, root.Value); {
		case result == 0:
			return root
		case result < 0:
			root = root.Left
		default:
			root = root.Right
		}
	}
	return nil
}

// InOrder returns an in-order traversal of the binary search tree rooted at root.
//
// For a tree built with Insert or InsertFunc, InOrder returns values in sorted order according
// to the tree's ordering.
// Time complexity: O(n). Additional space: O(n).
func InOrder[T any](root *BSTNode[T]) []T {
	out := make([]T, 0)
	stack := make([]*BSTNode[T], 0)
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, current.Value)
		current = current.Right
	}

	return out
}
