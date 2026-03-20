package graphs

// UnionFind keeps track of disjoint sets with path compression and union by rank.
type UnionFind struct {
	parent []int
	rank   []uint8
	count  int
}

// NewUnionFind allocates a disjoint-set structure for nodes in the range [0, size).
func NewUnionFind(size int) *UnionFind {
	if size < 0 {
		size = 0
	}
	parent := make([]int, size)
	rank := make([]uint8, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank, count: size}
}

// Find returns the representative for node or -1 when node is out of range.
func (uf *UnionFind) Find(node int) int {
	if uf == nil || !validNode(node, len(uf.parent)) {
		return -1
	}
	root := node
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	for node != root {
		parent := uf.parent[node]
		uf.parent[node] = root
		node = parent
	}
	return root
}

// Union merges the sets containing a and b and reports whether a merge happened.
func (uf *UnionFind) Union(a, b int) bool {
	if uf == nil {
		return false
	}
	rootA := uf.Find(a)
	rootB := uf.Find(b)
	if rootA == -1 || rootB == -1 || rootA == rootB {
		return false
	}
	if uf.rank[rootA] < uf.rank[rootB] {
		rootA, rootB = rootB, rootA
	}
	uf.parent[rootB] = rootA
	if uf.rank[rootA] == uf.rank[rootB] {
		uf.rank[rootA]++
	}
	uf.count--
	return true
}

// Connected reports whether a and b belong to the same set.
func (uf *UnionFind) Connected(a, b int) bool {
	if uf == nil {
		return false
	}
	rootA := uf.Find(a)
	return rootA != -1 && rootA == uf.Find(b)
}

// Count returns the number of remaining disjoint sets.
func (uf *UnionFind) Count() int {
	if uf == nil {
		return 0
	}
	return uf.count
}
