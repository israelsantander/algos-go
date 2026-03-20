# github.com/israelsantander/algos-go

`github.com/israelsantander/algos-go` is a standalone Go algorithms library built around reusable, plain Go APIs.

Target Go version: `1.26.1`

## Warning

This repository is for learning, experimentation, and reference.
Do not use it as-is in production systems.

## Goals

- Provide pure algorithm implementations with direct function APIs.
- Keep mutation semantics explicit.
- Offer a lightweight optional catalog for discovery and docs.

## Packages

- `sorting`: bubble, selection, insertion, shell, merge, quick, heap, counting, radix, reverse, quickselect
- `searching`: linear, binary, exponential, jump, BST, AVL
- `graphs`: DFS, BFS, topological sort, Dijkstra, Bellman-Ford, Prim, Kruskal, Tarjan, cycle detection, connected components, union-find, plus input adapters
- `linear`: generic stack and queue
- `lists`: singly, doubly, circular, circular doubly linked lists
- `recursion`: Towers of Hanoi, factorial, Fibonacci, permutations, combinations, subsets, N-Queens
- `catalog`: metadata-only algorithm catalog

## API Conventions

- Sorting functions expose both copy-returning and in-place variants.
- Sorting APIs are generic:
  - ordered variants such as `sorting.Bubble` and `sorting.QuickInPlace` work with `cmp.Ordered` types
  - `...Func` variants accept a comparator for custom types and custom orderings
- integer-specific sorting APIs such as `sorting.Counting` and `sorting.Radix` target `[]int`
- selection is exposed separately through `sorting.QuickSelect`
- Search functions return plain values or small tree/result structs.
- Graph functions use explicit input formats and result structs where needed.
- Data structure packages provide reusable types and methods.

## Sorting Examples

Ordered values:

```go
values := []int{5, 1, 4, 2, 8}
sorted := sorting.Bubble(values)
sorting.QuickInPlace(values)
```

Custom types with a comparator:

```go
type person struct {
	Name string
	Age  int
}

people := []person{
	{Name: "Ana", Age: 29},
	{Name: "Bo", Age: 18},
	{Name: "Eve", Age: 24},
}

sorted := sorting.QuickFunc(people, func(a, b person) bool {
	return a.Age < b.Age
})
```

## Status

This project is focused on clean reusable APIs, examples, benchmarks, and tests for studying algorithms in Go.

## Tests

Run the full test suite with:

```bash
make test
```

Test targets are verbose by default, so you see individual test names as they run.

Run tests for a single package:

```bash
make test-pkg PKG=./sorting
make test-sorting
make test-searching
make test-lists
make test-graphs
make test-linear
make test-recursion
```

Pass extra `go test` arguments after `--`:

```bash
make test-sorting -- -run TestReverse
make test -- -count=1 -failfast
```

## Benchmarks

Run the full benchmark suite with:

```bash
make bench
```

Run focused benchmarks by package family with:

```bash
make bench-sorting
make bench-searching
make bench-graphs
```

Run a focused benchmark pattern for one package:

```bash
make bench-pkg PKG=./sorting BENCH=Quick COUNT=3
make bench-sorting BENCH=BubbleInPlace COUNT=1
```

The benchmark suite is intentionally data-rich:

- sorting benchmarks vary by input size and pattern (`random`, `sorted`, `reversed`, `nearly-sorted`)
- searching benchmarks vary by input size and target position (`first`, `middle`, `last`, `missing`)
- graph benchmarks vary by graph size and density (`sparse`, `dense`)
- data structure and recursion benchmarks exercise realistic operation sequences rather than single calls
