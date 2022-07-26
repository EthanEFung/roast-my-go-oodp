package main

type publisher[T any] interface {
  subscribe(*subscriber[T])
  unsubscribe(*subscriber[T])
  notify()
}
