package main

import "time"

func main() {
	/*
		in our example, we are creating the functionality of a tram, and providing methods
		to operate on it's different parts. A tram needs a way to open its doors so that
		passengers can board, the doors need to close to ensure safe travel, it needs
		needs to move and stop etc.
	*/
	tram := newTram()
	t := time.NewTicker(time.Second)
	var moving bool
	for range t.C {
		tram.openDoors()
		tram.closeDoors()
		tram.requestStop()
		if moving {
			tram.move()
			tram.stop()
		} else {
			tram.stop()
			tram.move()
		}
		moving = !moving
		// However on the interval we want to make sure that some of these methods don't execute.
		// see `tram.go`.
	}

}
