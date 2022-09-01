package main

import "fmt"

/*
	strategyLRU is a concrete cache strategy
*/
type strategyLRU struct {
}

func (s strategyLRU) execute(c *cache) {
	fmt.Println("executing least recently used eviction")
}
