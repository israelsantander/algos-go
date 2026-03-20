package graphs

// BellmanFord computes single-source shortest paths on a weighted adjacency matrix.
//
// A zero weight means "no edge". Unlike Dijkstra, BellmanFord supports negative edge weights.
// It returns an error when the graph contains a negative-weight cycle reachable from start.
func BellmanFord(matrix [][]int, start int) (BellmanFordResult, error) {
	n := len(matrix)
	if !validNode(start, n) {
		return BellmanFordResult{}, errInvalidStart
	}

	distances := make([]int, n)
	parents := make([]int, n)
	seen := make([]bool, n)
	for i := range distances {
		distances[i] = infDistance
		parents[i] = -1
	}
	distances[start] = 0
	seen[start] = true
	order := make([]int, 0, n)

	for iteration := 0; iteration < n-1; iteration++ {
		changed := false
		for from, row := range matrix {
			if distances[from] == infDistance {
				continue
			}
			for to, weight := range row {
				if weight == 0 {
					continue
				}
				candidate := distances[from] + weight
				if candidate < distances[to] {
					distances[to] = candidate
					parents[to] = from
					if !seen[to] {
						seen[to] = true
						order = append(order, to)
					}
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

	for from, row := range matrix {
		if distances[from] == infDistance {
			continue
		}
		for to, weight := range row {
			if weight == 0 {
				continue
			}
			if distances[from]+weight < distances[to] {
				return BellmanFordResult{}, errNegativeCycle
			}
		}
	}

	for i, value := range distances {
		if value == infDistance {
			distances[i] = -1
		}
	}
	return BellmanFordResult{Distances: distances, Parents: parents, Order: order}, nil
}
