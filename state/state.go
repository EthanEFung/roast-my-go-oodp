package main

type state interface {
	move()
	stop()
	closeDoors()
	openDoors()
	requestStop() 
}