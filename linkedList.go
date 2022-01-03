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
	myList.deleteByValue(15)
	myList.printList()
	myList.insertAfter(100, 456)
	myList.printList()
	myList.deleteLinkedList()
	myList.printList()
}

type node struct {
	next *node
	data int
}

type linkedList struct {
	head   *node
	length int
}

//Insert in the front of the Linked List
func (l *linkedList) insertBefore(value int) {
	newNode := node{data: value}

	if l.head != nil { //   45>98>14
		newNode.next = l.head //21>45>98>14
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
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

//Delete a node by its value
func (l *linkedList) deleteByValue(value int) {
	if l.head == nil { //or l.length==0
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--

}

//Delete the linked list except head
func (l *linkedList) deleteLinkedList() {
	if l.head == nil { //or l.length==0
		return
	}
	currentNode := l.head
	currentNode.next = nil

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
