package main

import "fmt"

/*
	television is a concrete receiver struct that implements the device interface
	and stores the "running" state and methods to change its state
*/
type television struct {
	running bool
}

func (device *television) on() {
	device.running = true
	fmt.Println("television is on")
}
func (device *television) off() {
	device.running = false
	fmt.Println("television is off")
}
