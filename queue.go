package main

import "fmt"

type queue struct {
	data []int
}

// enqueue
func (q *queue) enqueue(val1 int) {
	q.data = append(q.data, val1)
}

// dequeue

func (q *queue) dequeue() {
	q.data = q.data[1:]
}

func (q *queue) print() {
	fmt.Println(q.data)
}

func main() {

	myQ := queue{}
	myQ.enqueue(22)
	myQ.enqueue(23)
	myQ.enqueue(24)
	myQ.enqueue(25)
	myQ.print()
	myQ.dequeue()
	myQ.dequeue()
	myQ.print()
}
