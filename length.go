package main

import "fmt"

func main() {
	fmt.Println("Hlo")
	myList := linkedList{}
	myList.insertBefore(21)
	myList.insertBefore(15)
	myList.insertBefore(456)
	myList.insertBefore(1895)
	myList.printList()
	myList.insertAfter(100, 456)
	myList.printList()
	l := myList.lengthLinkedList()
	fmt.Println(l)

}

type node struct {
	next *node
	data int
}

type linkedList struct {
	head *node
}

//Insert in the front of the Linked List
func (l *linkedList) insertBefore(value int) {
	newNode := node{data: value}

	if l.head != nil { //   45>98>14
		newNode.next = l.head //21>45>98>14
		l.head = &newNode

	} else {
		l.head = &newNode

	}
	return
}

//Insert after a specific node of the Linked List
func (l *linkedList) insertAfter(value, after int) {
	newNode := node{data: value}
	insertAfter := l.head

	for insertAfter != nil {
		if insertAfter.data == after {
			newNode.next = insertAfter.next
			insertAfter.next = &newNode

		}
		insertAfter = insertAfter.next
	}

}

func (l *linkedList) lengthLinkedList() int {
	currentNode := l.head
	var length int
	for currentNode != nil {
		length++
		currentNode = currentNode.next

	}
	return length
}

//Traveersing through and printing the list
func (l *linkedList) printList() {
	if l.head == nil { //or l.length==0
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}
