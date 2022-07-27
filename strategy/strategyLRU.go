package main

import "fmt"

type strategyLRU struct {
}

func (s strategyLRU) execute(c *cache) {
	fmt.Println("executing least recently used eviction")
}
