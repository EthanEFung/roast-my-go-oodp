package main

func main() {
	lru := strategyLRU{}
	lfu := strategyLFU{}
	fifo := strategyFIFO{}
	c := newCache(lru)
	c.add("k1", "v1")
	c.add("k2","v2")
	c.add("k3", "v3")

	c.setStrategy(lfu)
	c.add("k4", "v4")
	
	c.setStrategy(fifo)
	c.add("k5", "v5")
}