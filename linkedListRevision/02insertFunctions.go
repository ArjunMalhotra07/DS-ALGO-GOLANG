package main

func (listReceiver *list) insertBeginning(incomingValue string) {
	newNode := node{value: incomingValue}

	if listReceiver.head == nil {
		listReceiver.head = &newNode
		listReceiver.tail = &newNode
		listReceiver.length++
	} else {
		listReceiver.head.prev = &newNode
		newNode.next = listReceiver.head
		listReceiver.head = &newNode
		listReceiver.length++
	}

}

func (listReceiver *list) insertBefore(insertBeforeString string, incomingValue string) {
	newNode := node{value: incomingValue}
	currentNode := listReceiver.head
	for currentNode.next != nil {

		if currentNode.next.value == insertBeforeString {
			currentNode.next.prev = &newNode
			newNode.next = currentNode.next
			newNode.prev = currentNode
			currentNode.next = &newNode

			break
		}
		currentNode = currentNode.next
	}

}

func (listReceiver *list) insertAfter(insertAfterString string, incomingValue string) {
	newNode := node{value: incomingValue}

	currentNode := listReceiver.head

	for currentNode.next != nil {
		if currentNode.value == insertAfterString {
			currentNode.next.prev = &newNode
			newNode.next = currentNode.next
			newNode.prev = currentNode
			currentNode.next = &newNode
			break
		}
		currentNode = currentNode.next
	}

}

func (listReceiver *list) insertLast(incomingValue string) {
	newNode := node{value: incomingValue}

	currentNode := listReceiver.head

	if listReceiver.head == nil {
		listReceiver.head = &newNode
		listReceiver.tail = &newNode

		listReceiver.length++
	} else {
		for currentNode.next != nil {

			currentNode = currentNode.next
		}
		listReceiver.tail = &newNode
		currentNode.next = &newNode
		newNode.next = nil
		newNode.prev = currentNode
		listReceiver.length++
	}

}
