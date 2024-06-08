package main

type queueNode struct {
	value piastrella
	next  *queueNode
}

type queue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func (q *queue) Enqueue(v piastrella) {
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

func (q *queue) Dequeue() (piastrella, bool) {
	if q.Len() == 0 {
		return piastrella{}, false
	}
	v := q.head.value
	q.head = q.head.next
	q.length--
	return v, true
}

func (q *queue) Len() int {
	return q.length
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}

/*
func (q *queue) isEmpty() bool {
	return q.head == nil
}
*/
