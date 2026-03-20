package graphs

// DFS performs a depth-first traversal from start and returns visit order.
//
// It returns nil when start is out of range.
func DFS(adj [][]int, start int) []int {
	if start < 0 || start >= len(adj) {
		return nil
	}
	visited := make([]bool, len(adj))
	order := make([]int, 0, len(adj))
	stack := []int{start}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node < 0 || node >= len(adj) || visited[node] {
			continue
		}
		visited[node] = true
		order = append(order, node)
		for i := len(adj[node]) - 1; i >= 0; i-- {
			next := adj[node][i]
			if next >= 0 && next < len(adj) && !visited[next] {
				stack = append(stack, next)
			}
		}
	}
	return order
}

// BFS performs a breadth-first traversal from start and returns visit order.
//
// It returns nil when start is out of range.
func BFS(adj [][]int, start int) []int {
	if start < 0 || start >= len(adj) {
		return nil
	}
	visited := make([]bool, len(adj))
	visited[start] = true
	queue := []int{start}
	order := make([]int, 0, len(adj))
	for head := 0; head < len(queue); head++ {
		node := queue[head]
		order = append(order, node)
		for _, next := range adj[node] {
			if next >= 0 && next < len(adj) && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return order
}
