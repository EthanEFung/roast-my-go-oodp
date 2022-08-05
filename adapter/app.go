package main

import (
	"fmt"
)

type app struct {
	config string
}

func (client app) setJSONConfig(bytes []byte) {
	client.config = string(bytes)
	fmt.Println("config", client.config)
}