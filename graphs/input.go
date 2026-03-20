package graphs

// AdjacencyList builds an adjacency-list representation from edges.
//
// Edges with out-of-range endpoints are ignored. For undirected graphs, each edge is added
// in both directions.
func AdjacencyList(nodeCount int, edges []Edge, directed bool) [][]int {
	adj := make([][]int, nodeCount)
	for _, edge := range edges {
		if edge.From < 0 || edge.From >= nodeCount || edge.To < 0 || edge.To >= nodeCount {
			continue
		}
		adj[edge.From] = append(adj[edge.From], edge.To)
		if !directed {
			adj[edge.To] = append(adj[edge.To], edge.From)
		}
	}
	return adj
}

// WeightedMatrix builds an adjacency-matrix representation from edges.
//
// Edges with out-of-range endpoints are ignored. A zero edge weight is normalized to one so zero
// can continue to mean "no edge" inside the matrix representation.
func WeightedMatrix(nodeCount int, edges []Edge, directed bool) [][]int {
	matrix := make([][]int, nodeCount)
	for i := range matrix {
		matrix[i] = make([]int, nodeCount)
	}
	for _, edge := range edges {
		if edge.From < 0 || edge.From >= nodeCount || edge.To < 0 || edge.To >= nodeCount {
			continue
		}
		weight := edge.Weight
		if weight == 0 {
			weight = 1
		}
		matrix[edge.From][edge.To] = weight
		if !directed {
			matrix[edge.To][edge.From] = weight
		}
	}
	return matrix
}
