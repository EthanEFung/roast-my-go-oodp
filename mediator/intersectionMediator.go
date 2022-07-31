package main

type intersectionMediator struct {
	lightQueue []*trafficLight
}

func newIntersection() *intersectionMediator {
	return &intersectionMediator{
		lightQueue: make([]*trafficLight, 0),
	}
}

func (m *intersectionMediator) service() {
	for {
		first := m.lightQueue[0]
		first.call() // here the mediator issues a message to the traffic light
		m.lightQueue = append(m.lightQueue[1:], first)
	}
}

/*
	requestCrossing is an example of a message passed to the mediator
	which will be passed to the traffic light
*/
func (m *intersectionMediator) requestCrossing(light *trafficLight) {
	light.requestCrossing()
}

/*
	append is not apart of the pattern, just utility function
	that makes it easier to create the lightQueue
*/
func (m *intersectionMediator) append(light *trafficLight) {
	m.lightQueue = append(m.lightQueue, light)
}
