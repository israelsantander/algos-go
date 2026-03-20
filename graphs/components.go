package graphs

// HasCycleDirected reports whether adj contains a directed cycle.
func HasCycleDirected(adj [][]int) bool {
	const (
		unseen = iota
		active
		done
	)

	type frame struct {
		node int
		next int
	}

	state := make([]uint8, len(adj))
	stack := make([]frame, 0, len(adj))
	for start := range adj {
		if state[start] != unseen {
			continue
		}
		stack = append(stack[:0], frame{node: start})
		for len(stack) > 0 {
			top := &stack[len(stack)-1]
			if state[top.node] == unseen {
				state[top.node] = active
			}
			if top.next == len(adj[top.node]) {
				state[top.node] = done
				stack = stack[:len(stack)-1]
				continue
			}
			next := adj[top.node][top.next]
			top.next++
			if !validNode(next, len(adj)) {
				continue
			}
			switch state[next] {
			case unseen:
				stack = append(stack, frame{node: next})
			case active:
				return true
			}
		}
	}
	return false
}

// HasCycleUndirected reports whether adj contains an undirected cycle.
func HasCycleUndirected(adj [][]int) bool {
	uf := NewUnionFind(len(adj))
	seen := make(map[[2]int]struct{})
	for from, neighbors := range adj {
		for _, to := range neighbors {
			if !validNode(to, len(adj)) {
				continue
			}
			key := [2]int{from, to}
			if to < from {
				key = [2]int{to, from}
			}
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			if uf.Connected(from, to) {
				return true
			}
			uf.Union(from, to)
		}
	}
	return false
}

// ConnectedComponents returns the connected components of an undirected graph.
func ConnectedComponents(adj [][]int) [][]int {
	visited := make([]bool, len(adj))
	queue := make([]int, 0, len(adj))
	components := make([][]int, 0)
	for start := range adj {
		if visited[start] {
			continue
		}
		component := make([]int, 0)
		queue = append(queue[:0], start)
		visited[start] = true
		for head := 0; head < len(queue); head++ {
			node := queue[head]
			component = append(component, node)
			for _, next := range adj[node] {
				if validNode(next, len(adj)) && !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		components = append(components, component)
	}
	return components
}
