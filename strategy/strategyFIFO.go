package main

import (
	"fmt"
)

type strategyFIFO struct {
}

func (s strategyFIFO) execute(c *cache) {
	fmt.Println("executing first in first out eviction")
}
