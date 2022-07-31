package main

type crosswalk struct {
	mediator *intersectionMediator
	light    *trafficLight
}

func (c *crosswalk) request() {
	/*
		  here we see that event though the crosswalk has
			reference to the light in which they would like to
			request a crossing from, this function still
			passes its message through the mediator
	*/
	c.mediator.requestCrossing(c.light)
}
