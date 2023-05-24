package queue

import "testing"

func BenchmarkSlice(b *testing.B) {
	var q []int
	for i := 0; i < 64; i++ {
		q = append(q, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q = q[1:]
		q = append(q, i)
	}
}

func BenchmarkSliceWithGrow(b *testing.B) {
	var q []int
	for i := 0; i < b.N; i++ {
		q = append(q, i)
		q = append(q, i)
		q = q[1:]
		if i%1024 == 0 {
			q = nil
		}
	}
}
