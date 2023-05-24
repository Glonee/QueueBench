package queue

import (
	"testing"
)

func BenchmarkRingBuffer(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < 64; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
		q.Push(i)
	}
}

func BenchmarkRingBufferWithGrow(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Push(i)
		q.Pop()
		if i%1024 == 0 {
			q = RingBuffer[int]{}
		}
	}
}

func BenchmarkRingBufferChecked(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < 64; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PopCheckAndRwrite()
		q.Push(i)
	}
}

func BenchmarkRingBufferCheckedWithGrow(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Push(i)
		q.PopCheckAndRwrite()
		if i%1024 == 0 {
			q = RingBuffer[int]{}
		}
	}
}
