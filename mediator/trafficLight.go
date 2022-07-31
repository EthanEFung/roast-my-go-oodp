package main

import (
	"fmt"
	"time"
)

type trafficLight struct {
	id         string
	enableWalk bool
	mediator   *intersectionMediator
}

func (c *trafficLight) call() {
	ticker := time.NewTicker(time.Second * 2)
	fmt.Printf("LIGHT %v --> Green\n", c.id)
	for i := 0; i < 5; i++ {
		<-ticker.C
		if i == 2 {
			fmt.Printf("LIGHT %v --> Yellow\n", c.id)
		}
		/*
			Ideally the code run proceeding this comment
			would make a request to mediator which would then
			hand the message to the crosswalk. In other words,
			here is an opportunity to refactor into the
			mediator pattern further.
		*/
		if c.enableWalk && i != 4 {
			fmt.Printf("CROSS %v --> Walk\n", c.id)
		} else {
			fmt.Printf("CROSS %v --> Wait\n", c.id)
		}
	}
	c.enableWalk = false
	fmt.Printf("LIGHT %v --> Red\n", c.id)
}
func (c *trafficLight) requestCrossing() {
	/*
		here, we see that the traffic light has
		a method to communicate an intention to
		make a request to crosswalk for a future call
	*/
	c.enableWalk = true
}
