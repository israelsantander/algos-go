package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/israelsantander/algos-go/catalog"
	"github.com/israelsantander/algos-go/graphs"
	"github.com/israelsantander/algos-go/recursion"
	"github.com/israelsantander/algos-go/searching"
	"github.com/israelsantander/algos-go/sorting"
)

type demoFunc func(io.Writer) error

var demoRegistry = map[string]demoFunc{
	"avl":          demoAVL,
	"bellman-ford": demoBellmanFord,
	"binary":       demoBinary,
	"bubble":       demoBubble,
	"counting":     demoCounting,
	"dfs":          demoDFS,
	"dijkstra":     demoDijkstra,
	"exponential":  demoExponential,
	"insertion":    demoInsertion,
	"kruskal":      demoKruskal,
	"nqueens":      demoNQueens,
	"quickselect":  demoQuickSelect,
	"selection":    demoSelection,
	"tarjan":       demoTarjan,
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		writeUsage(stdout)
		return 0
	}

	switch args[0] {
	case "help", "-h", "--help":
		writeUsage(stdout)
		return 0
	case "list":
		return runList(args[1:], stdout, stderr)
	case "show":
		return runShow(args[1:], stdout, stderr)
	case "demo":
		return runDemo(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n\n", args[0])
		writeUsage(stderr)
		return 1
	}
}

func runList(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	fs.SetOutput(stderr)

	var family string
	fs.StringVar(&family, "family", "", "filter entries by family")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if fs.NArg() != 0 {
		fmt.Fprintln(stderr, "list does not accept positional arguments")
		return 1
	}

	entries := catalog.Entries()
	if family != "" {
		if !validFamily(family) {
			fmt.Fprintf(stderr, "invalid family %q\n", family)
			return 1
		}
		entries = filterEntriesByFamily(entries, catalog.Family(family))
	}

	tw := tabwriter.NewWriter(stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tFAMILY\tNAME\tCOMPLEXITY")
	for _, entry := range entries {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", entry.ID, entry.Family, entry.Name, entry.Complexity)
	}
	if err := tw.Flush(); err != nil {
		fmt.Fprintf(stderr, "failed to write list output: %v\n", err)
		return 1
	}
	return 0
}

func runShow(args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "usage: algos-go show <id>")
		return 1
	}

	entry, ok := findEntry(args[0])
	if !ok {
		fmt.Fprintf(stderr, "unknown algorithm id %q\n", args[0])
		return 1
	}

	fmt.Fprintf(stdout, "Name: %s\n", entry.Name)
	fmt.Fprintf(stdout, "ID: %s\n", entry.ID)
	fmt.Fprintf(stdout, "Family: %s\n", entry.Family)
	fmt.Fprintf(stdout, "Complexity: %s\n", entry.Complexity)
	fmt.Fprintf(stdout, "Example: %s\n", entry.Example)
	return 0
}

func runDemo(args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "usage: algos-go demo <id>")
		return 1
	}

	entry, ok := findEntry(args[0])
	if !ok {
		fmt.Fprintf(stderr, "unknown algorithm id %q\n", args[0])
		return 1
	}

	demo, ok := demoRegistry[entry.ID]
	if !ok {
		fmt.Fprintf(stderr, "no demo registered for %q\n", entry.ID)
		return 1
	}

	fmt.Fprintf(stdout, "%s (%s)\n", entry.Name, entry.ID)
	if err := demo(stdout); err != nil {
		fmt.Fprintf(stderr, "demo failed for %q: %v\n", entry.ID, err)
		return 1
	}
	return 0
}

func writeUsage(w io.Writer) {
	fmt.Fprintln(w, "algos-go exposes the catalog and deterministic demos for this repository.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  algos-go list")
	fmt.Fprintln(w, "  algos-go list -family <sorting|searching|graphs|recursion>")
	fmt.Fprintln(w, "  algos-go show <id>")
	fmt.Fprintln(w, "  algos-go demo <id>")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Available families: sorting, searching, graphs, recursion")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  go run ./cmd/algos-go list")
	fmt.Fprintln(w, "  go run ./cmd/algos-go list -family sorting")
	fmt.Fprintln(w, "  go run ./cmd/algos-go show bubble")
	fmt.Fprintln(w, "  go run ./cmd/algos-go demo dijkstra")
}

func validFamily(value string) bool {
	switch catalog.Family(value) {
	case catalog.FamilySorting, catalog.FamilySearching, catalog.FamilyGraphs, catalog.FamilyRecursion:
		return true
	default:
		return false
	}
}

func filterEntriesByFamily(entries []catalog.Entry, family catalog.Family) []catalog.Entry {
	filtered := make([]catalog.Entry, 0, len(entries))
	for _, entry := range entries {
		if entry.Family == family {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func findEntry(id string) (catalog.Entry, bool) {
	for _, entry := range catalog.Entries() {
		if entry.ID == id {
			return entry, true
		}
	}
	return catalog.Entry{}, false
}

func demoBubble(w io.Writer) error {
	values := []int{5, 1, 4, 2, 8}
	fmt.Fprintf(w, "input: %v\n", values)
	fmt.Fprintf(w, "sorted: %v\n", sorting.Bubble(values))
	return nil
}

func demoCounting(w io.Writer) error {
	values := []int{4, -1, 2, -1}
	fmt.Fprintf(w, "input: %v\n", values)
	fmt.Fprintf(w, "sorted: %v\n", sorting.Counting(values))
	return nil
}

func demoQuickSelect(w io.Writer) error {
	values := []int{9, 1, 7, 3, 5}
	k := 2
	value, ok := sorting.QuickSelect(values, k)
	if !ok {
		return errors.New("quickselect returned no value")
	}
	fmt.Fprintf(w, "input: %v\n", values)
	fmt.Fprintf(w, "k: %d\n", k)
	fmt.Fprintf(w, "value: %d\n", value)
	return nil
}

func demoSelection(w io.Writer) error {
	values := []int{64, 25, 12, 22, 11}
	fmt.Fprintf(w, "input: %v\n", values)
	fmt.Fprintf(w, "sorted: %v\n", sorting.Selection(values))
	return nil
}

func demoInsertion(w io.Writer) error {
	values := []int{31, 41, 59, 26}
	fmt.Fprintf(w, "input: %v\n", values)
	fmt.Fprintf(w, "sorted: %v\n", sorting.Insertion(values))
	return nil
}

func demoBinary(w io.Writer) error {
	values := []int{1, 3, 5, 7, 9}
	target := 7
	fmt.Fprintf(w, "values: %v\n", values)
	fmt.Fprintf(w, "target: %d\n", target)
	fmt.Fprintf(w, "index: %d\n", searching.Binary(values, target))
	return nil
}

func demoExponential(w io.Writer) error {
	values := []int{1, 3, 5, 7, 9}
	target := 7
	fmt.Fprintf(w, "values: %v\n", values)
	fmt.Fprintf(w, "target: %d\n", target)
	fmt.Fprintf(w, "index: %d\n", searching.Exponential(values, target))
	return nil
}

func demoAVL(w io.Writer) error {
	var root *searching.AVLNode[int]
	for _, value := range []int{10, 5, 15, 7} {
		root = searching.AVLInsert(root, value)
	}
	found := searching.AVLSearch(root, 7)
	if found == nil {
		return errors.New("expected AVLSearch to find 7")
	}
	fmt.Fprintf(w, "inorder: %v\n", searching.AVLInOrder(root))
	fmt.Fprintf(w, "target: %d\n", 7)
	fmt.Fprintf(w, "found: %d\n", found.Value)
	return nil
}

func demoDFS(w io.Writer) error {
	adj := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}
	fmt.Fprintf(w, "adjacency: %v\n", adj)
	fmt.Fprintf(w, "start: %d\n", 0)
	fmt.Fprintf(w, "order: %v\n", graphs.DFS(adj, 0))
	return nil
}

func demoDijkstra(w io.Writer) error {
	matrix := [][]int{
		{0, 4, 1, 0},
		{0, 0, 2, 1},
		{0, 1, 0, 5},
		{0, 0, 0, 0},
	}
	result, err := graphs.Dijkstra(matrix, 0)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "start: %d\n", 0)
	fmt.Fprintf(w, "distances: %v\n", result.Distances)
	fmt.Fprintf(w, "parents: %v\n", result.Parents)
	fmt.Fprintf(w, "order: %v\n", result.Order)
	return nil
}

func demoBellmanFord(w io.Writer) error {
	matrix := [][]int{
		{0, 4, 5, 0},
		{0, 0, -2, 6},
		{0, 0, 0, 3},
		{0, 0, 0, 0},
	}
	result, err := graphs.BellmanFord(matrix, 0)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "start: %d\n", 0)
	fmt.Fprintf(w, "distances: %v\n", result.Distances)
	fmt.Fprintf(w, "parents: %v\n", result.Parents)
	fmt.Fprintf(w, "order: %v\n", result.Order)
	return nil
}

func demoKruskal(w io.Writer) error {
	edges := []graphs.Edge{
		{From: 0, To: 1, Weight: 10},
		{From: 0, To: 2, Weight: 6},
		{From: 0, To: 3, Weight: 5},
		{From: 1, To: 3, Weight: 15},
		{From: 2, To: 3, Weight: 4},
	}
	result, err := graphs.Kruskal(edges, 4)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "edges: %v\n", edges)
	fmt.Fprintf(w, "total-weight: %d\n", result.TotalWeight)
	fmt.Fprintf(w, "components: %d\n", result.Components)
	fmt.Fprintf(w, "forest: %v\n", result.Edges)
	return nil
}

func demoTarjan(w io.Writer) error {
	adj := [][]int{
		{1},
		{2, 3},
		{0},
		{4},
		{},
	}
	components := graphs.Tarjan(adj)
	for _, component := range components {
		slices.Sort(component)
	}
	slices.SortFunc(components, func(a, b []int) int {
		return strings.Compare(fmt.Sprint(a), fmt.Sprint(b))
	})
	fmt.Fprintf(w, "adjacency: %v\n", adj)
	fmt.Fprintf(w, "components: %v\n", components)
	return nil
}

func demoNQueens(w io.Writer) error {
	solutions := recursion.NQueens(4)
	fmt.Fprintf(w, "n: %d\n", 4)
	fmt.Fprintf(w, "solutions: %d\n", len(solutions))
	fmt.Fprintf(w, "placements: %v\n", solutions)
	return nil
}
