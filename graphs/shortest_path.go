package graphs

// Dijkstra computes single-source shortest paths on a weighted adjacency matrix.
//
// A zero weight means "no edge". Negative weights are rejected with an error.
func Dijkstra(matrix [][]int, start int) (DijkstraResult, error) {
	n := len(matrix)
	if !validNode(start, n) {
		return DijkstraResult{}, errInvalidStart
	}
	distances := make([]int, n)
	parents := make([]int, n)
	visited := make([]bool, n)
	for i := range distances {
		distances[i] = infDistance
		parents[i] = -1
	}
	distances[start] = 0
	order := make([]int, 0, n)

	for range n {
		node := -1
		best := infDistance
		for i := 0; i < n; i++ {
			if !visited[i] && distances[i] < best {
				best = distances[i]
				node = i
			}
		}
		if node == -1 {
			break
		}
		visited[node] = true
		order = append(order, node)
		for next, weight := range matrix[node] {
			if weight < 0 {
				return DijkstraResult{}, errNegativeWeights
			}
			if weight == 0 || visited[next] {
				continue
			}
			candidate := distances[node] + weight
			if candidate < distances[next] {
				distances[next] = candidate
				parents[next] = node
			}
		}
	}

	for i, value := range distances {
		if value == infDistance {
			distances[i] = -1
		}
	}
	return DijkstraResult{Distances: distances, Parents: parents, Order: order}, nil
}
