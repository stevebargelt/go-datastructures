package queue

// Queue is a FIFO data structure
type Queue struct {
	start, end *node
	length     int
}
type node struct {
	value interface{}
	next  *node
}

// New creates a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Dequeue returns the next item from the front of the queue
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return n.value
}

// Enqueue places an item at the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	n := &node{value, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

// Len returns the length of the queue
func (q *Queue) Len() int {
	return q.length
}

// Peek returns the first item off the queue without removing it from the queue
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}
