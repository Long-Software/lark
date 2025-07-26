package queue

type Queue[T any] struct {
	first *Node[T]
	last  *Node[T]
	len   int
}

func New[T any](val T) *Queue[T] {
	node := newNode(val)
	return &Queue[T]{
		first: node,
		last:  node,
		len:   1,
	}
}

func (q *Queue[T]) Enqueue(val T) {
	node := newNode(val)
	if q.IsEmpty() {
		q.first, q.last = node, node
	} else {
		q.last.next = node
		q.last = node
	}
	q.len++
}

func (q *Queue[T]) Dequeue() *Node[T] {
	if q.IsEmpty() {
		return nil
	}
	tmp := q.first
	if q.len == 1 {
		q.first, q.last = nil, nil
	} else {
		q.first = q.first.next
		tmp.next = nil
	}
	q.len--
	return tmp
}
func (q *Queue[T]) Length() int {
	return q.len
}
func (q *Queue[T]) First() *Node[T] {
	return q.first
}
func (q *Queue[T]) Last() *Node[T] {
	return q.last
}

func (q *Queue[T]) IsEmpty() bool {
	return q.len == 0
}
