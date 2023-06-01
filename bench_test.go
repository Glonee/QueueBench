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
	for i := 0; i < 2; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

func BenchmarkRingBuffer(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

func BenchmarkRingBufferBuffered(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < 2; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}
