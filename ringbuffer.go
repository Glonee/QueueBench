package queue

type RingBuffer[T any] struct {
	ring             []T
	headPos, tailPos int
	full             bool
}

func (r *RingBuffer[T]) Init(size int) {
	r.ring = make([]T, size)
}

func (r *RingBuffer[T]) Len() int {
	if r.full {
		return len(r.ring)
	}
	if r.tailPos >= r.headPos {
		return r.tailPos - r.headPos
	}
	return r.tailPos - r.headPos + len(r.ring)
}

func (r *RingBuffer[T]) Push(t T) {
	if r.full || len(r.ring) == 0 {
		r.grow()
	}
	r.ring[r.tailPos] = t
	r.tailPos++
	if r.tailPos == len(r.ring) {
		r.tailPos = 0
	}
	if r.tailPos == r.headPos {
		r.full = true
	}
}

func (r *RingBuffer[T]) Pop() T {
	if !r.full && r.headPos == r.tailPos {
		panic("github.com/quic-go/quic-go/internal/utils/ringbuffer: pop from a empty queue")
	}
	r.full = false
	t := r.ring[r.headPos]
	var zeroValue T
	r.ring[r.headPos] = zeroValue
	r.headPos++
	if r.headPos == len(r.ring) {
		r.headPos = 0
	}
	return t
}

// Grow the maximum size of the queue.
// This method assume the queue is full.
func (r *RingBuffer[T]) grow() {
	oldRing := r.ring
	newSize := len(oldRing) * 2
	if newSize == 0 {
		newSize = 1
	}
	r.ring = make([]T, newSize)
	headLen := copy(r.ring, oldRing[r.headPos:])
	copy(r.ring[headLen:], oldRing[:r.headPos])
	r.headPos, r.tailPos, r.full = 0, len(oldRing), false
}

func (r *RingBuffer[T]) Clear() {
	var zeroValue T
	for i := range r.ring {
		r.ring[i] = zeroValue
	}
	r.headPos, r.tailPos, r.full = 0, 0, false
}
