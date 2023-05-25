package queue

import (
	"testing"
)

func BenchmarkRingBuffer(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

// func BenchmarkRingBufferWithGrow(b *testing.B) {
// 	q := RingBuffer[int]{}
// 	for i := 0; i < b.N; i++ {
// 		q.Push(i)
// 		q.Push(i)
// 		q.Pop()
// 		if i%1024 == 0 {
// 			q = RingBuffer[int]{}
// 		}
// 	}
// }

func BenchmarkRingBufferWithCheck(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.PopWithCheckAndRwrite()
	}
}

// func BenchmarkRingBufferWithCheckAndGrow(b *testing.B) {
// 	q := RingBuffer[int]{}
// 	for i := 0; i < b.N; i++ {
// 		q.Push(i)
// 		q.Push(i)
// 		q.PopWithCheckAndRwrite()
// 		if i%1024 == 0 {
// 			q = RingBuffer[int]{}
// 		}
// 	}
// }
