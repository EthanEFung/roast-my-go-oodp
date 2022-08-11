package main

import (
	"fmt"
)

/*
app is a concrete client
*/
type app struct {
	config string
}

func (client app) exportJSON(bytes []byte) {
	client.config = string(bytes)
	fmt.Println("config", client.config)
}