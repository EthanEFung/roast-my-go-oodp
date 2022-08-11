# Adapter Pattern
The adapter pattern falls into the category of structural patterns that
focuses on creating integrations between two object interfaces that would've
otherwise been incompatible.

Generally, within this pattern, there is a `Client` object that is within the developers control with a method and `Service`
outside of the developers control with useful functionality, but is incompatible with
the client. Generally speaking this pattern can be useful if the developer does not want to change the client objects, but wants to extend the functionality of the application. This pattern focuses on creating concrete `Adapter`s that adhere to the
client's interface but holds a reference to the service.

To see the output, run the following cli commands from the terminal
```bash
go run .
```