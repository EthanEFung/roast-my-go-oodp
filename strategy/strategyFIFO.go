package main

import (
	"fmt"
)

/*
	strategyFIFO is a concrete cache strategy
*/
type strategyFIFO struct {
}

func (s strategyFIFO) execute(c *cache) {
	fmt.Println("executing first in first out eviction")
}
