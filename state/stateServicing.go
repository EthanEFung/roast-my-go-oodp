package main

import "fmt"

type stateServicing struct {
	tram *tram
}

func (state stateServicing) move() {
	fmt.Println("cannot move tram while in service")
}
func (state stateServicing) stop() {
	fmt.Println("tram already stopped")
}
func (state stateServicing) closeDoors() {
	fmt.Println("cannot close doors while servicing")
}
func (state stateServicing) openDoors() {
	fmt.Println("tram doors are already open")
}
func (state stateServicing) requestStop() {
	fmt.Println("tram cannot receive request while servicing")
} 