package Queues

import (
	"fmt"
)

const MAX_SIZE = 100

type Queue struct {
	items [MAX_SIZE]int
	FRONT int
	REAR  int
}

func (q *Queue) Enqueue(value int) {
	if q.REAR == MAX_SIZE-1 {
		fmt.Println("Queue is full")
		return
	} else if q.FRONT == -1 && q.REAR == -1 {
		q.FRONT = 0
		q.REAR = 0
	} else {
		q.REAR++
	}
	q.items[q.REAR] = value
}

func (q *Queue) Dequeue() {
	if q.FRONT == -1 {
		fmt.Println("Queue is empty")
		return
	} else if q.FRONT == q.REAR {
		q.FRONT = -1
		q.REAR = -1
	} else {
		q.FRONT++
	}
}

func (q *Queue) Display() {
	if q.FRONT == -1 {
		fmt.Println("Queue is empty")
		return
	}
	for i := q.FRONT; i <= q.REAR; i++ {
		fmt.Print(q.items[i], " ")
	}
	fmt.Println()
}

func Queues_main() {
	queue := Queue{FRONT: -1, REAR: -1}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	queue.Display()

	queue.Dequeue()
	queue.Display()
}
