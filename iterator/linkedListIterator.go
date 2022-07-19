package main

/* implements the iterator interface */
type linkedListIterator struct {
	current *listNode
}

func (i *linkedListIterator) next() int {
	if i.current == nil {
		return 0
	}
	returned := i.current
	i.current = i.current.next
	return returned.val
}

func (i *linkedListIterator) done() bool {
	return i.current == nil
}
