package main

import "fmt"

type cycleNodes struct {
	value string
	next  *cycleNodes
}

type cycleList struct {
	head   *cycleNodes
	length int
}

func callCycleFunction(cycleListObject cycleList) {
	cycleListObject.insertAtStart("Test 1")
	cycleListObject.insertAtStart("Test 2")
	cycleListObject.insertAtStart("Test 3")
	cycleListObject.insertAtStart("Test 4")
	cycleListObject.printList()
	cycleListObject.checkCycleExists()
	cycleListObject.makeCycle("Test 3")
	cycleListObject.printList()
	cycleListObject.checkCycleExists()
}

// 2 pointer approach
func (listReceiver *cycleList) checkCycleExists() {
	f := fmt.Println
	// At 1st node
	oneStepPointer := listReceiver.head
	// At 2nd Node
	twoStepPointer := listReceiver.head.next

	flag := false
	for oneStepPointer != twoStepPointer {
		oneStepPointer = oneStepPointer.next
		twoStepPointer = twoStepPointer.next.next
		// if they come at same position then circular LL
		if oneStepPointer == twoStepPointer {
			flag = true
			break
		}
		// if twoStepPointer or twoStepPointer.next reaches nil not a circular LL
		if twoStepPointer == nil || twoStepPointer.next == nil {
			flag = false
			break
		}
	}
	// Print Results
	if flag {
		f("Cycle exist")
	} else {
		f("No Cycles")
	}

}

// Connects tail to head
func (listReceiver *cycleList) makeCycle(incomingValue string) {

	currentNode := listReceiver.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = listReceiver.head

}

// inserts element at the start of the LL
func (listReceiver *cycleList) insertAtStart(incomingValue string) {
	newNode := cycleNodes{value: incomingValue}

	if listReceiver.head == nil {
		listReceiver.head = &newNode
		listReceiver.length++
	} else {
		newNode.next = listReceiver.head
		listReceiver.head = &newNode
		listReceiver.length++
	}

}

// Prints LL 2 times if circular and if no circle then only once
func (l *cycleList) printList() {
	count := 0
	f := fmt.Println
	headNode := l.head
	f()
	f("Linked List From Front -- ")

	for headNode != nil {
		if headNode.value == "Test 4" {
			count++
		}
		if count == 3 {
			break
		}
		fmt.Printf("->%s", headNode.value)
		headNode = headNode.next

	}
	f()

}
