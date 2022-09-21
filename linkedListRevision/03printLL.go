package main

import "fmt"

func (l *list) printList() {
	printNewLine := fmt.Println
	headNode := l.head
	printNewLine()
	printNewLine("Linked List From Front -- ")

	for headNode != nil {
		fmt.Printf("->%s", headNode.value)
		headNode = headNode.next
	}
	printNewLine()
	printNewLine()
	lastNode := l.tail

	printNewLine("Linked List From Tail -- ")

	for lastNode != nil {
		fmt.Printf("->%s", lastNode.value)
		lastNode = lastNode.prev
	}
	printNewLine()
}
