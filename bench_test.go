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

func BenchmarkRingBuffer(b *testing.B) {
	q := RingBuffer[int]{}
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}

func BenchmarkRingBufferWithCheck(b *testing.B) {
	q := RingBuffer[int]{}
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
