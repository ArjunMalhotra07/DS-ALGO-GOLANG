package main

import "fmt"

func main() {
	f := fmt.Println
	myQueue := Queue{}
	myQueue.Enque(100)
	myQueue.Enque(200)
	myQueue.Enque(300)
	f(myQueue)
	f("Dequed", myQueue.Deque())

	f(myQueue)
}

type Queue struct {
	items []int
}

func (q *Queue) Enque(i int) {
	q.items = append(q.items, i)
}
func (q *Queue) Deque() int {
	l := len(q.items)
	toRemove := q.items[0]

	q.items = q.items[1:l]
	return toRemove
}
