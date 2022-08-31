package main

/*
	state is the interface that represent all the methods that are possible for the
	tram object
*/
type state interface {
	move()
	stop()
	closeDoors()
	openDoors()
	requestStop()
}
