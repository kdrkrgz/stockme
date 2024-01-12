package types

type Queue []any

func NewQueue(size int) *Queue {
	q := make(Queue, size)
	return &q
}

func (q *Queue) Push(element any) {
	// push element to tail
	*q = append(*q, element)
}

func (q *Queue) Pop() any {
	// pop element from head and return element
	rq := *q
	var el any
	el, *q = rq[0], rq[1:]
	return el
}

func (q *Queue) RightShift(element any) {
	// inital queue object cap equals len and contains nil elements
	// when RightShift called first nil elements removed and add given element to tail
	q.Pop()
	q.Push(element)
}

func (q *Queue) Elements() any {
	return *q
}
