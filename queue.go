package main

type queueNode struct {
	value int
	next  *queueNode
}

type queue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func (q *queue) Enqueue(v int) {
	newNode := &queueNode{v, nil}
	if q.Len() == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.length++
}

func (q *queue) Dequeue() (int, bool) {
	if q.Len() == 0 {
		return 0, false
	}
	v := q.head.value
	q.head = q.head.next
	q.length--
	return v, true
}

func (q *queue) Len() int {
	return q.length
}
