# mediator

mediator is a behaviour design pattern that attempts to reduce the dependencies
between objects by reducing the multi-faceted object communication of an
application through a mediator.  An issue that could come from this approach is as
more and more objects send messages through a mediator, the more likely the mediator
object will become bloated. However, having a singular place that messages are past
to can help with overall readability if the size of the mediator is reasonably small.
Visit the `main.go` file to learn more and run
```
go run .
```
to see the output of the main function. The program is continuous so make sure to
terminate the process.