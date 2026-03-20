package linear

import (
	"fmt"
	"testing"
)

func BenchmarkStackPushPop(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		b.Run(fmt.Sprintf("StackPushPop/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var s Stack[int]
				for j := 0; j < size; j++ {
					s.Push(j)
				}
				for j := 0; j < size; j++ {
					_, _ = s.Pop()
				}
			}
		})
	}
}

func BenchmarkQueueEnqueueDequeue(b *testing.B) {
	for _, size := range []int{32, 256, 4096, 65536} {
		b.Run(fmt.Sprintf("QueueEnqueueDequeue/n=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var q Queue[int]
				for j := 0; j < size; j++ {
					q.Enqueue(j)
				}
				for j := 0; j < size; j++ {
					_, _ = q.Dequeue()
				}
			}
		})
	}
}
