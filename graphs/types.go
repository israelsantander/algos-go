package graphs

// Edge describes a weighted edge between two graph nodes.
type Edge struct {
	From   int
	To     int
	Weight int
}

// DijkstraResult contains the shortest-path data computed by Dijkstra.
type DijkstraResult struct {
	Distances []int
	Parents   []int
	Order     []int
}

// BellmanFordResult contains the shortest-path data computed by BellmanFord.
type BellmanFordResult struct {
	Distances []int
	Parents   []int
	Order     []int
}

// PrimResult contains the spanning-tree data computed by Prim.
type PrimResult struct {
	TotalWeight int
	Parents     []int
	Keys        []int
	Order       []int
	Edges       []Edge
}

// KruskalResult contains the spanning-forest data computed by Kruskal.
type KruskalResult struct {
	TotalWeight int
	Edges       []Edge
	Components  int
}
