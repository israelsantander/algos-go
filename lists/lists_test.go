package lists

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	var list SinglyLinkedList[int]
	list.Append(1)
	list.Append(3)
	if !list.InsertAt(1, 2) {
		t.Fatal("insert failed")
	}
	if got := list.Values(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
	if value, ok := list.DeleteAt(1); !ok || value != 2 {
		t.Fatalf("delete got %d ok=%v", value, ok)
	}
}

func TestDoublyLinkedList(t *testing.T) {
	var list DoublyLinkedList[int]
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	if got := list.Values(); !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Fatalf("got %v", got)
	}
	if value, ok := list.DeleteAt(1); !ok || value != 1 {
		t.Fatalf("delete got %d ok=%v", value, ok)
	}
}

func TestDoublyLinkedListDeleteFromTailSide(t *testing.T) {
	var list DoublyLinkedList[int]
	for _, value := range []int{0, 1, 2, 3, 4} {
		list.Append(value)
	}
	if value, ok := list.DeleteAt(4); !ok || value != 4 {
		t.Fatalf("delete got %d ok=%v", value, ok)
	}
	if got := list.Values(); !reflect.DeepEqual(got, []int{0, 1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestCircularLists(t *testing.T) {
	var singly CircularLinkedList[int]
	singly.Append(1)
	singly.Append(2)
	singly.Append(3)
	if got := singly.Values(3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}

	var doubly CircularDoublyLinkedList[int]
	doubly.Append(4)
	doubly.Append(5)
	doubly.Append(6)
	if got := doubly.Values(3); !reflect.DeepEqual(got, []int{4, 5, 6}) {
		t.Fatalf("got %v", got)
	}
}
