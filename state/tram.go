package main

/*
	tram is the context in this pattern which stores a reference to a current state
*/
type tram struct {
	state state
	moving *stateMoving
	stopped *stateStopped
	servicing *stateServicing
}

func newTram() *tram {
	result := new(tram)
	result.moving = &stateMoving{result}
	result.stopped = &stateStopped{result}
	result.servicing = &stateServicing{result}
	result.state = result.stopped
	return result
}

func (context *tram) setState(state state) {
	context.state = state
}

func (context *tram) move() {
	context.state.move()
}
func (context *tram) stop() {
	context.state.stop()
}
func (context *tram) closeDoors() {
	context.state.closeDoors()
}
func (context *tram) openDoors() {
	context.state.openDoors()
}
func (context *tram) requestStop() {
	context.state.requestStop()
} 