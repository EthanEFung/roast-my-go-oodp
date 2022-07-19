package main

import "fmt"

func main() {
	head := &listNode{1, nil}
	list := linkedList{head, head}

	a, b := &listNode{2, nil}, &listNode{3, nil}
	list.queue(a)
	list.queue(b)

	iter := list.createIterator()

	for !iter.done() {
		val := iter.next()
		fmt.Println(val)
	}
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
