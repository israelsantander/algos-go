package recursion

// Permutations returns every permutation of values in traversal order.
func Permutations[T any](values []T) [][]T {
	if values == nil {
		return nil
	}
	current := cloneSlice(values)
	out := make([][]T, 0)
	var generate func(int)
	generate = func(index int) {
		if index == len(current) {
			out = append(out, cloneSlice(current))
			return
		}
		for i := index; i < len(current); i++ {
			current[index], current[i] = current[i], current[index]
			generate(index + 1)
			current[index], current[i] = current[i], current[index]
		}
	}
	generate(0)
	return out
}

// Combinations returns every size-k combination of values in traversal order.
func Combinations[T any](values []T, k int) [][]T {
	switch {
	case values == nil:
		return nil
	case k < 0 || k > len(values):
		return nil
	case k == 0:
		return [][]T{{}}
	}

	out := make([][]T, 0)
	current := make([]T, 0, k)
	var choose func(int)
	choose = func(start int) {
		if len(current) == k {
			out = append(out, cloneSlice(current))
			return
		}
		remaining := k - len(current)
		for i := start; i <= len(values)-remaining; i++ {
			current = append(current, values[i])
			choose(i + 1)
			current = current[:len(current)-1]
		}
	}
	choose(0)
	return out
}

// Subsets returns the power set of values in traversal order.
func Subsets[T any](values []T) [][]T {
	if values == nil {
		return nil
	}
	out := make([][]T, 0, 1<<len(values))
	current := make([]T, 0, len(values))
	var build func(int)
	build = func(index int) {
		if index == len(values) {
			out = append(out, cloneSlice(current))
			return
		}
		build(index + 1)
		current = append(current, values[index])
		build(index + 1)
		current = current[:len(current)-1]
	}
	build(0)
	return out
}

func cloneSlice[T any](values []T) []T {
	if values == nil {
		return nil
	}
	out := make([]T, len(values))
	copy(out, values)
	return out
}
