package main

func main() {
	// in our example here, the television is the receiver. See `device.go`
	tv := &television{}

	// typically we ill have multiple concrete commands. See `command.go` first.
	on := &onCommand{
		device: tv, // notice that the commands store reference to the receiver
	}

	off := &offCommand{
		device: tv,
	}

	// each command is seperated from these concrete invokers
	onButton := &button{
		command: on, // each concrete invoker will store a reference to a concrete command
	}

	offButton := &button{
		command: off,
	}

	/*
		Now, under the hood each invoker executes the command that was assigned to them.
		Each command, having reference to the receiver object will be able to execute business
		logic and interact with the receiver.
	*/
	onButton.click()
	offButton.click()
}
