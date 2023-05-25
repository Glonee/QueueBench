package queue

import (
	"testing"
)

func BenchmarkQueue(b *testing.B) {
	q := Queue[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

// func BenchmarkQueueWithGrow(b *testing.B) {
// 	q := Queue[int]{}
// 	for i := 0; i < b.N; i++ {
// 		q.Push(i)
// 		q.Push(i)
// 		q.Pop()
// 		if i%1024 == 0 {
// 			q = Queue[int]{}
// 		}
// 	}
// }
