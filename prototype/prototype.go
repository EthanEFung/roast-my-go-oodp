package main

/*
  prototype is an interface that contains the ability to "clone"
	or copy the existing data
*/
type prototype[entity any] interface {
	clone() entity
}