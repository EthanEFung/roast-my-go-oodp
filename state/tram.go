package main

/*
	tram is the context in this pattern which stores a reference to a current state
*/
type tram struct {
	state     state // here we see that we store reference to a `state`
	moving    *stateMoving
	stopped   *stateStopped
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

/*
	setState is what is used for transitions from one state to the next
*/
func (context *tram) setState(state state) {
	context.state = state
}

/*
	here can see that all of the functionality that is being used is a function of the
	trams `state`
*/

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
