# command

The command pattern is a behavioral design pattern that focuses on seperating
the algorithms of the application from the objects that call them. Typically, an
`invoker` object will receive a concrete `command` object, and the concrete commands
are objects that store reference to a `receiver` object. This pattern can be useful
when the command is repeatedly used by different invokers, or when an invoker
is used for many different commands. Visit the `main.go` file to get started.
From the command line run
```
go run .
```
to see the output executed