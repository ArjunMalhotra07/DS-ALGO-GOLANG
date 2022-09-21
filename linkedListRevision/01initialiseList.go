package main

func initialise() {
	linkedListObject := list{}
	newListObject := testList{}
	cycleListObject := cycleList{}
	veriousFunctions(linkedListObject, 1, newListObject)
	callCycleFunction(cycleListObject)

}

type node struct {
	value string
	next  *node
	prev  *node
}

type list struct {
	head   *node
	tail   *node
	length int
}

func veriousFunctions(object list, index int, object2 testList) {
	object.insertBeginning("10")
	object.insertBeginning("00")
	object.insertBeginning("01")
	object.insertBeginning("02")
	object.insertLast("10")
	object.insertLast("00")
	object.insertLast("01")
	object.insertLast("02")

	// object.insertBefore("Arjun", "Nani")
	// object.insertAfter("Nani", "US")
	// object.insertLast("Pink")
	object.printList()
	// object.reverse()
	// object.printList()

	object.checkPalindrome(object2)

}
