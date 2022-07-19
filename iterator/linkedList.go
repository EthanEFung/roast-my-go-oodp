package main

/* implements the collection interface */
type linkedList struct {
	head    *listNode
	tail *listNode
}

type listNode struct {
	val  int
	next *listNode
}

func (collection *linkedList) queue(node *listNode) {
	collection.tail.next = node
	collection.tail = collection.tail.next
}

func (collection *linkedList) dequeue() *listNode {
	if collection.head == collection.tail {
		collection.tail = collection.tail.next
	}
	result := collection.head
	collection.head = collection.head.next
	return result
}

func (collection *linkedList) createIterator() iterator {
	return &linkedListIterator{collection.head}
}
