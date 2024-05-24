package main

type queue struct {
	items []int
}

func (q *queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

func (q *queue) Dequeue() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}
