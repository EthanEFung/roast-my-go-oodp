package main

/*
dress is a flyweight object that contains
immutable data
*/
type dress interface {
	getColor() string
}