# factory method

factory method is a creational pattern that seperates the concrete product from the
creation process. Go does not have support for inheritance which is generally used for
this pattern. However, in Go we can create a factory struct that can produce products
with a common interface and utilize the interface for the concrete products. Start with
the `main.go` file to get started, and from the terminal run 
```
go run .
```
