/*
  mediator is a behaviour design pattern that attempts to
	reduce the dependencies between objects by reducing
	the multi-faceted object communication of an application
	through a mediator

	a major issue that could come from this approach is as
	more and more objects send messages through a mediator,
	the more likely the mediator object will become bloated.
*/
package main

import "time"

func main() {
	intersection := newIntersection()
	lightA := &trafficLight{"A", false, intersection}
	lightB := &trafficLight{"B", false, intersection}

	/*
		  ** tangent **
			for these crosswalk components, this implementation
			is passing the pointers of the lights to indicate to
			the mediator which light the request is for. That said
			this could be an opportunity to pass an id of sorts which
			would have its own benefits, but imo fails to
			"reduce dependencies"
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
