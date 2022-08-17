package main

import "fmt"

/*
	button is a concrete sender (invoker)
*/
type button struct {
	command command
}

func (sender *button) click() {
	fmt.Println(sender.command.name(), "button click")
	sender.command.execute()
}
