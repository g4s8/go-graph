package search

type intQueue struct {
	first *intQNode
	last  *intQNode
}

type intQNode struct {
	val  int
	next *intQNode
}

func (q *intQueue) enq(val int) {
	node := &intQNode{val: val}
	if q.first == nil {
		q.first = node
		q.last = node
		return
	}
	q.last.next = node
	q.last = node
}

func (q *intQueue) deq() int {
	if q.first == nil {
		panic("queue is empty")
	}
	node := q.first
	q.first = node.next
	if q.first == nil {
		q.last = nil
	}
	return node.val
}

func (q *intQueue) empty() bool {
	return q.first == nil
}

func (q *intQueue) len() int {
	var l int
	for node := q.first; node != nil; node = node.next {
		l++
	}
	return l
}

type intStack struct {
	head *intStackNode
	size int
}

type intStackNode struct {
	val  int
	next *intStackNode
}

func (s *intStack) push(val int) {
	head := s.head
	s.head = &intStackNode{val: val, next: head}
	s.size++
}

func (s *intStack) pushAll(vals []int) {
	for _, v := range vals {
		s.push(v)
	}
}

func (s *intStack) pop() int {
	if s.head == nil {
		panic("empty stack")
	}
	head := s.head
	s.head = head.next
	s.size--
	return head.val
}

func (s *intStack) empty() bool {
	return s.head == nil
}

func (s *intStack) slice() []int {
	return s.sliceReverted()
}

func (s *intStack) sliceReverted() []int {
	res := make([]int, s.size)
	for head, pos := s.head, s.size-1; head != nil; head, pos = head.next, pos-1 {
		res[pos] = head.val
	}
	return res
}
