package main

/*
	onCommand is a concrete command that turns
	a device "on"
*/
type onCommand struct {
	device device
}
func(command *onCommand) name() string {
	return "onCommand"
}
func(command *onCommand) execute() {
	command.device.on()
}
