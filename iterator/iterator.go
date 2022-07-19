package main

type iterator interface {
	next() int
	done() bool
}
