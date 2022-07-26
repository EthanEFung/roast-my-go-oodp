package main

type subscriber[C any] interface {
  update(C)
}
