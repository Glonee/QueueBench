package queue

import (
	"testing"
)

func BenchmarkQueue(b *testing.B) {
	q := Queue[int]{}
	for i := 0; i < 64; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
		q.Push(i)
	}
}

func BenchmarkQueueWithGrow(b *testing.B) {
	q := Queue[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Push(i)
		q.Pop()
		if i%1024 == 0 {
			q = Queue[int]{}
		}
	}
}
