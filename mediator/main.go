package main

import "time"

func main() {
	// in our example our mediator will be the intersection controller
	intersection := newIntersection()
	/*
		all other objects will hold a pointer to the intersection and pass messages to the
		mediator which will orchestrate processes
	*/
	lightA := &trafficLight{"A", false, intersection}
	lightB := &trafficLight{"B", false, intersection}

	/*
		## tangent ##
		for these crosswalk components, this implementation is passing the pointers of the
		lights to indicate to the mediator which light the request is for. That said this
		could be an opportunity to pass an id of sorts which would have its own benefits,
		but imo fails to "reduce dependencies"
	*/
	crosswalkA := crosswalk{intersection, lightA}
	crosswalkB := crosswalk{intersection, lightB}

	intersection.append(lightA)
	intersection.append(lightB)

	go intersection.service()

	time.Sleep(time.Second * 2)
	crosswalkB.request()
	time.Sleep(time.Second * 14)
	crosswalkA.request()
	crosswalkB.request()
	<-make(chan bool)
}
