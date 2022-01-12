package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hlo")
	myList := linkedList{}
	myList.insertBefore(21546)
	myList.insertBefore(21)
	myList.insertBefore(15)
	myList.insertBefore(456)
	myList.insertBefore(1895)
	myList.printList()
	myList.reverseLinkedList()
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

func (l *linkedList) reverseLinkedList() *node {
	var previousNode, nextNode *node
	currentNode := l.head
	if currentNode == nil {
		return currentNode
	}
	for currentNode != nil {
		nextNode, currentNode.next = currentNode.next, previousNode

		previousNode, currentNode = currentNode, nextNode

	}
	l.head = previousNode
	return l.head
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
