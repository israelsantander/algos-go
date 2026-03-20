package linear

import "testing"

func TestStack(t *testing.T) {
	var s Stack[int]
	if _, ok := s.Pop(); ok {
		t.Fatal("expected empty pop to fail")
	}
	s.Push(1)
	s.Push(2)
	if got, ok := s.Peek(); !ok || got != 2 {
		t.Fatalf("peek got %d ok=%v", got, ok)
	}
	if got, ok := s.Pop(); !ok || got != 2 {
		t.Fatalf("pop got %d ok=%v", got, ok)
	}
	if s.Len() != 1 {
		t.Fatalf("len got %d", s.Len())
	}
}

func TestQueue(t *testing.T) {
	var q Queue[int]
	if _, ok := q.Dequeue(); ok {
		t.Fatal("expected empty dequeue to fail")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	if got, ok := q.Peek(); !ok || got != 1 {
		t.Fatalf("peek got %d ok=%v", got, ok)
	}
	if got, ok := q.Dequeue(); !ok || got != 1 {
		t.Fatalf("dequeue got %d ok=%v", got, ok)
	}
	if q.Len() != 1 {
		t.Fatalf("len got %d", q.Len())
	}
}

func TestQueueValuesAfterCompaction(t *testing.T) {
	var q Queue[int]
	for i := 0; i < 8; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		if _, ok := q.Dequeue(); !ok {
			t.Fatalf("dequeue %d failed", i)
		}
	}
	if got, ok := q.Peek(); !ok || got != 5 {
		t.Fatalf("peek got %d ok=%v", got, ok)
	}
	if got := q.Values(); len(got) != 3 || got[0] != 5 || got[2] != 7 {
		t.Fatalf("values got %v", got)
	}
}
