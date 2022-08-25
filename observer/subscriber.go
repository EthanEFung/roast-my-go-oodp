package main

/*
  subscriber is an object in the observer pattern that exposes to
  the publisher a method that indicates an event has occured.
*/
type subscriber[C any] interface {
	update(C)
}
