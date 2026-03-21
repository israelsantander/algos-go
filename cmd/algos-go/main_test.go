package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/israelsantander/algos-go/catalog"
)

func TestRunHelpWhenNoArgs(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run(nil, &stdout, &stderr)
	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
	if !strings.Contains(stdout.String(), "Usage:") {
		t.Fatalf("stdout = %q, want usage output", stdout.String())
	}
}

func TestRunHelpCommand(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"help"}, &stdout, &stderr)
	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
	if !strings.Contains(stdout.String(), "go run ./cmd/algos-go demo dijkstra") {
		t.Fatalf("stdout = %q, want example command", stdout.String())
	}
}

func TestRunList(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"list"}, &stdout, &stderr)
	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}

	output := stdout.String()
	if !strings.Contains(output, "ID") || !strings.Contains(output, "COMPLEXITY") {
		t.Fatalf("stdout = %q, want table header", output)
	}
	if !strings.Contains(output, "bubble") || !strings.Contains(output, "nqueens") {
		t.Fatalf("stdout = %q, want catalog rows", output)
	}
}

func TestRunListWithFamily(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"list", "-family", "sorting"}, &stdout, &stderr)
	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}

	output := stdout.String()
	if !strings.Contains(output, "bubble") {
		t.Fatalf("stdout = %q, want sorting entries", output)
	}
	if strings.Contains(output, "binary") {
		t.Fatalf("stdout = %q, want only sorting entries", output)
	}
}

func TestRunListWithInvalidFamily(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"list", "-family", "invalid"}, &stdout, &stderr)
	if exitCode == 0 {
		t.Fatal("run() exit code = 0, want non-zero")
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), `invalid family "invalid"`) {
		t.Fatalf("stderr = %q, want invalid family error", stderr.String())
	}
}

func TestRunShow(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"show", "bubble"}, &stdout, &stderr)
	if exitCode != 0 {
		t.Fatalf("run() exit code = %d, want 0", exitCode)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}

	output := stdout.String()
	if !strings.Contains(output, "Name: Bubble Sort") {
		t.Fatalf("stdout = %q, want bubble details", output)
	}
	if !strings.Contains(output, "Example: sorting.Bubble([]int{5,1,4,2,8})") {
		t.Fatalf("stdout = %q, want example line", output)
	}
}

func TestRunShowMissingID(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"show", "missing-id"}, &stdout, &stderr)
	if exitCode == 0 {
		t.Fatal("run() exit code = 0, want non-zero")
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q, want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), `unknown algorithm id "missing-id"`) {
		t.Fatalf("stderr = %q, want unknown id error", stderr.String())
	}
}

func TestRunDemoRepresentativeFamilies(t *testing.T) {
	testCases := []struct {
		name      string
		id        string
		wantLines []string
	}{
		{
			name:      "sorting",
			id:        "bubble",
			wantLines: []string{"Bubble Sort (bubble)", "input: [5 1 4 2 8]", "sorted: [1 2 4 5 8]"},
		},
		{
			name:      "searching",
			id:        "binary",
			wantLines: []string{"Binary Search (binary)", "target: 7", "index: 3"},
		},
		{
			name:      "graphs",
			id:        "dijkstra",
			wantLines: []string{"Dijkstra (dijkstra)", "distances:", "parents:", "order:"},
		},
		{
			name:      "recursion",
			id:        "nqueens",
			wantLines: []string{"N-Queens (nqueens)", "solutions: 2", "placements: [[1 3 0 2] [2 0 3 1]]"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run([]string{"demo", tc.id}, &stdout, &stderr)
			if exitCode != 0 {
				t.Fatalf("run() exit code = %d, want 0", exitCode)
			}
			if stderr.Len() != 0 {
				t.Fatalf("stderr = %q, want empty", stderr.String())
			}

			output := stdout.String()
			for _, want := range tc.wantLines {
				if !strings.Contains(output, want) {
					t.Fatalf("stdout = %q, want line containing %q", output, want)
				}
			}
		})
	}
}

func TestDemoRegistryCoversCatalog(t *testing.T) {
	for _, entry := range catalog.Entries() {
		if _, ok := demoRegistry[entry.ID]; !ok {
			t.Fatalf("missing demo for catalog id %q", entry.ID)
		}
	}
}
