package main

import "fmt"

/*
adapter holds a reference to a service object,
yet adheres to the client interface
*/
type adapter struct {
	service service
	config string
}

func (a adapter) exportJSON([]byte) {
	a.config = a.service.exportYAML()

	// here we write conversions from the yaml -> json and export the json

	fmt.Println("adapted config", a.config)
}