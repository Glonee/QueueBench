package queue

type RingBuffer[T any] struct {
	zeroVal    T
	ring       []T
	head, tail int
	full       bool
}

func (r *RingBuffer[T]) Len() int {
	if r.full {
		return len(r.ring)
	}
	if r.head <= r.tail {
		return r.tail - r.head
	}
	return r.tail - r.head + len(r.ring)
}

func (r *RingBuffer[T]) grow() {
	oldRing := r.ring
	newSize := len(oldRing) * 2
	if newSize == 0 {
		newSize = 1
	}
	r.ring = make([]T, newSize)
	headLen := copy(r.ring, oldRing[r.head:])
	copy(r.ring[headLen:], oldRing[:r.head])
	r.head, r.tail, r.full = 0, len(oldRing), false
}

func (r *RingBuffer[T]) Push(t T) {
	if r.full || len(r.ring) == 0 {
		r.grow()
	}
	r.ring[r.tail] = t
	r.tail++
	if r.tail == len(r.ring) {
		r.tail = 0
	}
	if r.tail == r.head {
		r.full = true
	}
}

func (r *RingBuffer[T]) Pop() T {
	r.full = false
	t := r.ring[r.head]
	r.head++
	if r.head == len(r.ring) {
		r.head = 0
	}
	return t
}

func (r *RingBuffer[T]) PopCheckAndRwrite() (t T, ok bool) {
	if r.Len() == 0 {
		return
	}
	ok = true
	r.full = false
	t = r.ring[r.head]
	r.ring[r.head] = r.zeroVal
	r.head++
	if r.head == len(r.ring) {
		r.head = 0
	}
	return
}
