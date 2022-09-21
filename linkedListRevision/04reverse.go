package main

func (listReceiver *list) reverse() {
	currentNode := listReceiver.head
	var previousNode, nextNode *node
	listReceiver.tail = currentNode
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = previousNode
		currentNode.prev = nextNode
		previousNode = currentNode
		currentNode = nextNode
	}
	listReceiver.head = previousNode
}
