# proxy

proxy is a structural design pattern that focuses on interfacing with substitute
object instead of an object directly. The advantage of this pattern is layering
processes to run either before or after requests made to the original object. See
the `main.go` file to learn more. The example creates an http server locally on port
3000. From the terminal run
```
go run .
```
to see the data navigate to `localhost:3000` to see the data or from another terminal
window run

```
curl localhost:3000
```
