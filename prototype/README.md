# prototype

prototype is a creational pattern of copying existing objects. In some of the
more traditional oodp languages prototyping was necessary if attributes of a class
were private. In Go, all attributes of a struct are accessible. However, the pattern
can be useful when the initialization of an object is verbose, or the object itself
has a great deal of complexity making it hard to copy. Visit the `main.go` file and
from the terminal run
```
go run .
```
 