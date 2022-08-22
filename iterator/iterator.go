package main

/*
	iterator is the interface that no matter the underlying data structure of the
	collection should be exposed in our application if the collection is iterable.
*/
type iterator interface {
	next() int
	done() bool
}
