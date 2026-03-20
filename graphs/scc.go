package graphs

// Tarjan returns the strongly connected components of a directed graph.
func Tarjan(adj [][]int) [][]int {
	index := 0
	indexes := make([]int, len(adj))
	lowlink := make([]int, len(adj))
	onStack := make([]bool, len(adj))
	for i := range indexes {
		indexes[i] = -1
	}

	stack := make([]int, 0, len(adj))
	components := make([][]int, 0)

	var connect func(int)
	connect = func(node int) {
		indexes[node] = index
		lowlink[node] = index
		index++
		stack = append(stack, node)
		onStack[node] = true

		for _, next := range adj[node] {
			if !validNode(next, len(adj)) {
				continue
			}
			if indexes[next] == -1 {
				connect(next)
				if lowlink[next] < lowlink[node] {
					lowlink[node] = lowlink[next]
				}
				continue
			}
			if onStack[next] && indexes[next] < lowlink[node] {
				lowlink[node] = indexes[next]
			}
		}

		if lowlink[node] != indexes[node] {
			return
		}
		component := make([]int, 0)
		for {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			onStack[last] = false
			component = append(component, last)
			if last == node {
				break
			}
		}
		components = append(components, component)
	}

	for node := range adj {
		if indexes[node] == -1 {
			connect(node)
		}
	}
	return components
}
