package main

/*
service is an interface that only has an exportYAML function.
It is incompatible with the client interface.
*/
type service interface {
	exportYAML() string
}