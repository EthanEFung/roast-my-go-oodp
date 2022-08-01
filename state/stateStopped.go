package main

import "fmt"

type stateStopped struct {
	tram *tram
}

func (state stateStopped) move() {
	fmt.Println("moving departure")
	state.tram.setState(state.tram.moving)
}
func (state stateStopped) stop() {
	fmt.Println("tram already stopped")
}
func (state stateStopped) closeDoors() {
	fmt.Println("closing doors")
}
func (state stateStopped) openDoors() {
	fmt.Println("opening doors")
}
func (state stateStopped) requestStop() {
	fmt.Println("tram cannot receive requests to stop while stopped")
} 