package main

func main() {
	tv := &television{}

	on := &onCommand{
		device: tv,
	}

	off := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: on,
	}

	offButton := &button{
		command: off,
	}

	onButton.click()
	offButton.click()
}