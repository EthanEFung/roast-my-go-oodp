package main

import "fmt"

type stateMoving struct {
	tram *tram
}

func (state stateMoving) move() {
	fmt.Println("tram already in motion")
}
func (state stateMoving) stop() {
	fmt.Println("stopping")
	state.tram.setState(state.tram.stopped)
}
func (state stateMoving) closeDoors() {
	fmt.Println("doors already closed")
}
func (state stateMoving) openDoors() {
	fmt.Println("cannot open doors while in motion")
}
func (state stateMoving) requestStop() {
	fmt.Println("requesting stop")
} 