package main

type strategy[context any] interface {
	execute(context)
}
