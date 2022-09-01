package main

/*
	strategy is the object that
*/
type strategy[context any] interface {
	execute(context)
}
