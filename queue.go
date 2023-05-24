package queue

type Queue[T any] struct {
	head, tail []T
	headPos    int
}

func (q *Queue[T]) Len() int {
	return len(q.tail) - q.headPos + len(q.head)
}

func (q *Queue[T]) Push(t T) {
	q.tail = append(q.tail, t)
}

func (q *Queue[T]) Pop() T {
	if q.headPos >= len(q.head) {
		q.head, q.tail, q.headPos = q.tail, q.head[:0], 0
	}
	t := q.head[q.headPos]
	q.headPos++
	return t
}
