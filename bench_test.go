package queue

import (
	"testing"
)

func BenchmarkTwoStagesQueue(b *testing.B) {
	q := Queue[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

func BenchmarkTwoStagesQueueBuffered(b *testing.B) {
	q := Queue[int]{}
	for i := 0; i < 8; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

// func BenchmarkRingBuffer(b *testing.B) {
// 	q := RingBuffer[int]{}
// 	for i := 0; i < b.N; i++ {
// 		q.Push(i)
// 		q.Pop()
// 	}
// }

//	func BenchmarkRingBufferBuffered(b *testing.B) {
//		q := RingBuffer[int]{}
//		for i := 0; i < 8; i++ {
//			q.Push(i)
//		}
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			q.Push(i)
//			q.Pop()
//		}
//	}
func BenchmarkRingBufferWithCheck(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.PopWithCheckAndRwrite()
	}
}

func BenchmarkRingBufferWithCheckBuffered(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < 8; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.PopWithCheckAndRwrite()
	}
}
func BenchmarkSlice(b *testing.B) {
	var q []int
	for i := 0; i < b.N; i++ {
		q = append(q, i)
		q = q[1:]
	}
}

func BenchmarkSliceBuffered(b *testing.B) {
	var q []int
	for i := 0; i < 8; i++ {
		q = append(q, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q = append(q, i)
		q = q[1:]
	}
}
