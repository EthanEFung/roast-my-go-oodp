package main

type printable interface {
	print(indentation int) string
}