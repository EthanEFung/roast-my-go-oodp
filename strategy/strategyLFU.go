package main

import "fmt"

type strategyLFU struct {
}

func (s strategyLFU) execute(c *cache) {
	fmt.Println("executing least frequently used eviction")
}
