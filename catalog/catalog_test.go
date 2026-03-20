package catalog

import "testing"

func TestEntries(t *testing.T) {
	entries := Entries()
	if len(entries) == 0 {
		t.Fatal("expected catalog entries")
	}
	for _, entry := range entries {
		if entry.ID == "" || entry.Name == "" || entry.Family == "" {
			t.Fatalf("invalid entry: %+v", entry)
		}
	}
}
