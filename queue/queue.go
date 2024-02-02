package queue

type Queue[E any] interface {
	Add(e *E)
	Poll() *E
	Peek() *E
}

type LinkedQueue[E any] struct {
	buf []*E
}

func NewLinkedQueue[E any]() *LinkedQueue[E] {
	return &LinkedQueue[E]{
		buf: make([]*E, 0),
	}
}

func (q *LinkedQueue[E]) Add(e *E) {
	q.buf = append(q.buf, e)
}

func (q *LinkedQueue[E]) Poll() *E {
	element := q.buf[0]
	q.buf = q.buf[1:]
	return element
}

func (q *LinkedQueue[E]) Peek() *E {
	return q.buf[0]
}
