package main

type composite interface {
	search(string) bool
	getName() string
}