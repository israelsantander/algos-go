package graphs

// TopologicalSort returns a valid topological ordering for a directed acyclic graph.
//
// It returns an error when the graph contains a cycle.
func TopologicalSort(adj [][]int) ([]int, error) {
	indegree := make([]int, len(adj))
	for from := range adj {
		for _, to := range adj[from] {
			if to < 0 || to >= len(adj) {
				continue
			}
			indegree[to]++
		}
	}
	queue := make([]int, 0)
	for node, degree := range indegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}
	order := make([]int, 0, len(adj))
	for head := 0; head < len(queue); head++ {
		node := queue[head]
		order = append(order, node)
		for _, next := range adj[node] {
			if next < 0 || next >= len(adj) {
				continue
			}
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if len(order) != len(adj) {
		return nil, errGraphContainsCycle
	}
	return order, nil
}
