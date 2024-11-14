package search

const initQstackSize = 256

// intQstack is a queue and stack implementation using a slice.
type intQstack struct {
	a          []int
	head, tail int
}

// enq adds a value to the queue.
func (q *intQstack) enq(val int) {
	q.grow(1)
	q.a[q.tail] = val
	q.tail++
}

func (q *intQstack) deq() int {
	if q.head == q.tail {
		panic("queue is empty")
	}

	v := q.a[q.head]
	q.head++
	return v
}

func (q *intQstack) push(val int) {
	q.enq(val)
}

func (q *intQstack) pushAll(a []int) {
	q.grow(len(a))
	for _, v := range a {
		q.a[q.tail] = v
		q.tail++
	}
}

func (q *intQstack) pop() int {
	if q.head == q.tail {
		panic("stack is empty")
	}

	q.tail--
	v := q.a[q.tail]
	return v
}

func (q *intQstack) empty() bool {
	return q.head == q.tail
}

func (q *intQstack) len() int {
	return q.tail - q.head
}

func (q *intQstack) slice() []int {
	return q.a[q.head:q.tail]
}

func (q *intQstack) sliceReverted() []int {
	s := q.slice()
	res := make([]int, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}

func (q *intQstack) grow(n int) {
	if q.tail+n <= len(q.a) {
		return
	}

	if q.head > len(q.a)/2 {
		copy(q.a, q.a[q.head:])
		q.tail -= q.head
		q.head = 0
	}

	if q.tail+n > len(q.a) {
		newSize := 2 * len(q.a)
		if newSize < q.tail+n {
			newSize = q.tail + n
		}
		if newSize < initQstackSize {
			newSize = initQstackSize
		}
		newA := make([]int, newSize)
		copy(newA, q.a[q.head:q.tail])
		q.tail -= q.head
		q.head = 0
		q.a = newA
	}
}
