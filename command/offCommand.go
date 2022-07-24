package main

/*
	offCommand is a concrete command that turns a device "off"
*/
type offCommand struct {
	device device
}
func(command *offCommand) name() string {
	return "offCommand"
}
func(command *offCommand) execute() {
	command.device.off()
}
