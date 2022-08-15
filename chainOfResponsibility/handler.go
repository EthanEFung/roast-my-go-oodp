package main

/*
  handler is the interface that comprises of a method containing business logic,
  and another that sets a pointer or reference to the next handler in the chain
*/
type handler interface {
	execute(*order)
	setNext(handler)
}
