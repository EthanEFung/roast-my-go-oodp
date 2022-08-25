package main

/*
  publisher is an object in the observer pattern that has methods
  to add and remove subscribers from an internal collection, and a method
  to notify all subscribers of a specified event.
*/
type publisher[T any] interface {
	subscribe(*subscriber[T])
	unsubscribe(*subscriber[T])
	notify()
}
