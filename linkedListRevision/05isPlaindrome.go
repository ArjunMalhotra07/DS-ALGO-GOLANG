package main

// Make a new Linked List. Copy first half of the LL in it. sequentially check if LL matches the remaining ll.
import "fmt"

type testNode struct {
	value string
	next  *testNode
}

type testList struct {
	head   *testNode
	length int
}

func (listObject *list) checkPalindrome(incomingObject testList) {
	f := fmt.Println
	// Adding values to our LL
	currentNode := listObject.head
	for i := 0; i < (listObject.length)/2; i++ {
		incomingObject.insertBeginningofList(currentNode.value)
		currentNode = currentNode.next
	}

	// If length of LL is odd change the current node to next
	if (listObject.length)%2 != 0 {
		currentNode = currentNode.next
	}

	// Check top element of our LL with next Node value of the Linked List, pop head (change reference)
	flag := true
	newCurrent := incomingObject.head
	for newCurrent != nil {
		if currentNode.value != incomingObject.head.value {
			flag = false
		}
		currentNode = currentNode.next
		newCurrent = newCurrent.next
		incomingObject.head = newCurrent
	}
	// Print results
	if flag {
		f()
		f("Palindrome")
	} else {
		f("Not a Palindrome")
	}
}

func (listReceiver *testList) insertBeginningofList(incomingValue string) {
	newNode := testNode{value: incomingValue}

	if listReceiver.head == nil {
		listReceiver.head = &newNode
		listReceiver.length++
	} else {
		newNode.next = listReceiver.head
		listReceiver.head = &newNode
		listReceiver.length++
	}

}
