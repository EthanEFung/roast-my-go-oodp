/*
	In this example we hypothesize that we have a linked list data structure,
	but the implementation needs some fleshing out first. What we can do while
	we are developing this code is to first create an iterator interface,
	and as long as the iterator has the correct menthods, we can begin writing
	the framework of our algorithms that depend on our lists.
*/
package main

import "fmt"

func main() {
	head := &listNode{1, nil}
	list := linkedList{head, head}

	a, b := &listNode{2, nil}, &listNode{3, nil}
	list.queue(a)
	list.queue(b)

	/*
		above, we set up the data, and below we start
		instantiating different iterators
	*/
	iter := list.createIterator()

	for !iter.done() {
		// because we have a `next` method and a `done` method we have
		// what we need to traverse the collection
		val := iter.next()
		fmt.Println(val)
	}

	/*
		here we're setting up another list with a different iterator
	*/
	first := list.dequeue()
	nextList := linkedList{first, first}

	iter = nextList.createIterator()

	for !iter.done() {
		val := iter.next()
		fmt.Println(val)
	}

	first.next = nil // prevent infinite loop
	list.queue(first)

	iter = list.createIterator()
	for !iter.done() {
		val := iter.next()
		fmt.Println(val)
	}
}
