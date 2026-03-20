package graphs

// Prim computes a minimum spanning tree from start on a weighted adjacency matrix.
//
// A zero weight means "no edge". Negative weights are rejected with an error.
func Prim(matrix [][]int, start int) (PrimResult, error) {
	n := len(matrix)
	if !validNode(start, n) {
		return PrimResult{}, errInvalidStart
	}
	keys := make([]int, n)
	parents := make([]int, n)
	inTree := make([]bool, n)
	for i := range keys {
		keys[i] = infDistance
		parents[i] = -1
	}
	keys[start] = 0
	order := make([]int, 0, n)

	for range n {
		node := -1
		best := infDistance
		for i := 0; i < n; i++ {
			if !inTree[i] && keys[i] < best {
				best = keys[i]
				node = i
			}
		}
		if node == -1 {
			break
		}
		inTree[node] = true
		order = append(order, node)
		for next, weight := range matrix[node] {
			if weight < 0 {
				return PrimResult{}, errNegativeWeights
			}
			if weight == 0 || inTree[next] {
				continue
			}
			if weight < keys[next] {
				keys[next] = weight
				parents[next] = node
			}
		}
	}

	edges := make([]Edge, 0, n-1)
	total := 0
	for node, parent := range parents {
		if parent == -1 {
			continue
		}
		weight := matrix[parent][node]
		edges = append(edges, Edge{From: parent, To: node, Weight: weight})
		total += weight
	}
	for i, value := range keys {
		if value == infDistance {
			keys[i] = -1
		}
	}
	return PrimResult{TotalWeight: total, Parents: parents, Keys: keys, Order: order, Edges: edges}, nil
}
