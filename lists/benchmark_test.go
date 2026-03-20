package lists

import (
	"fmt"
	"testing"
)

func BenchmarkSinglyLinkedListAppend(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 16384} {
		b.Run(fmt.Sprintf("SinglyAppend/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var list SinglyLinkedList[int]
				for j := 0; j < size; j++ {
					list.Append(j)
				}
			}
		})
	}
}

func BenchmarkSinglyLinkedListDeleteMiddle(b *testing.B) {
	for _, size := range []int{32, 256, 4096} {
		b.Run(fmt.Sprintf("SinglyDeleteMiddle/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var list SinglyLinkedList[int]
				for j := 0; j < size; j++ {
					list.Append(j)
				}
				_, _ = list.DeleteAt(size / 2)
			}
		})
	}
}

func BenchmarkDoublyLinkedListAppend(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 16384} {
		b.Run(fmt.Sprintf("DoublyAppend/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var list DoublyLinkedList[int]
				for j := 0; j < size; j++ {
					list.Append(j)
				}
			}
		})
	}
}
