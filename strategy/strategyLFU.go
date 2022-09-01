package main

import "fmt"

/*
	strategyLFU is a concrete cache strategy
*/
type strategyLFU struct {
}

func (s strategyLFU) execute(c *cache) {
	fmt.Println("executing least frequently used eviction")
}
