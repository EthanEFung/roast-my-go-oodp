package main

import "time"

func main() {
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
	}
	
}