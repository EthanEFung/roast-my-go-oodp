package main

type handler interface {
	execute(*order)
	setNext(handler)
}
