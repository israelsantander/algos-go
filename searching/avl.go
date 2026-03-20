package searching

import "cmp"

// AVLNode is a node in an AVL tree.
//
// AVL trees rebalance after each insertion so lookup and insertion remain O(log n) even for
// sorted input.
type AVLNode[T any] struct {
	Value  T
	Left   *AVLNode[T]
	Right  *AVLNode[T]
	height int
}

// AVLInsert inserts value into an AVL tree and returns the root.
func AVLInsert[T cmp.Ordered](root *AVLNode[T], value T) *AVLNode[T] {
	return AVLInsertFunc(root, value, orderedCompare[T])
}

// AVLInsertFunc inserts value into an AVL tree using compare and returns the root.
//
// Duplicate values are inserted into the right subtree to keep behavior deterministic.
func AVLInsertFunc[T any](root *AVLNode[T], value T, compare func(a, b T) int) *AVLNode[T] {
	return avlInsert(root, value, requireCompare(compare))
}

func avlInsert[T any](root *AVLNode[T], value T, compare func(a, b T) int) *AVLNode[T] {
	if root == nil {
		return &AVLNode[T]{Value: value, height: 1}
	}
	if compare(value, root.Value) < 0 {
		root.Left = avlInsert(root.Left, value, compare)
	} else {
		root.Right = avlInsert(root.Right, value, compare)
	}
	updateAVLHeight(root)
	return rebalanceAVL(root)
}

// AVLSearch searches an AVL tree and returns the matching node or nil.
func AVLSearch[T cmp.Ordered](root *AVLNode[T], target T) *AVLNode[T] {
	return AVLSearchFunc(root, target, orderedCompare[T])
}

// AVLSearchFunc searches an AVL tree using compare and returns the matching node or nil.
func AVLSearchFunc[T any](root *AVLNode[T], target T, compare func(a, b T) int) *AVLNode[T] {
	compare = requireCompare(compare)
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

// AVLInOrder returns an in-order traversal of an AVL tree.
func AVLInOrder[T any](root *AVLNode[T]) []T {
	out := make([]T, 0)
	stack := make([]*AVLNode[T], 0)
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

func avlHeight[T any](node *AVLNode[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

func updateAVLHeight[T any](node *AVLNode[T]) {
	node.height = 1 + max(avlHeight(node.Left), avlHeight(node.Right))
}

func avlBalance[T any](node *AVLNode[T]) int {
	if node == nil {
		return 0
	}
	return avlHeight(node.Left) - avlHeight(node.Right)
}

func rotateAVLLeft[T any](node *AVLNode[T]) *AVLNode[T] {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	updateAVLHeight(node)
	updateAVLHeight(newRoot)
	return newRoot
}

func rotateAVLRight[T any](node *AVLNode[T]) *AVLNode[T] {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	updateAVLHeight(node)
	updateAVLHeight(newRoot)
	return newRoot
}

func rebalanceAVL[T any](node *AVLNode[T]) *AVLNode[T] {
	switch balance := avlBalance(node); {
	case balance > 1:
		if avlBalance(node.Left) < 0 {
			node.Left = rotateAVLLeft(node.Left)
		}
		return rotateAVLRight(node)
	case balance < -1:
		if avlBalance(node.Right) > 0 {
			node.Right = rotateAVLRight(node.Right)
		}
		return rotateAVLLeft(node)
	default:
		return node
	}
}
