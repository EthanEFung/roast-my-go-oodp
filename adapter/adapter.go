package main

import "fmt"

type adapter struct {
	service service
	config string
}

func(a adapter) convert() {
	a.config = a.service.exportYAML()
	// here we'd use something the translate the yaml file to json
}

func (a adapter) setJSONConfig([]byte) {
	fmt.Println("adapted config", a.config)
}