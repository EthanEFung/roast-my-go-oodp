# singleton

singleton is creational pattern that focuses on having one instance of an object with
global access. This pattern is used when it is necessary to limit or manage access to
a shared resource. In the example found in `main.go` we focus more on how to implement
a singleton over an example use case since it can be tricky to properly implement the
pattern when orchestrating go routines (which is often when the singleton would need
to be implemented. From the terminal run 
```
go run main.go
```
to see the output of the main function.
