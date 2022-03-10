package main

import "fmt"

func main() {
	fmt.Println("Hello")
	myQueue := Queue{}
	myQueue.Enque(100)
	myQueue.Enque(200)
	myQueue.Enque(300)
	fmt.Println(myQueue)
	removed := myQueue.Deque()
	fmt.Println(removed, "Dequed")
	fmt.Println(myQueue)
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
