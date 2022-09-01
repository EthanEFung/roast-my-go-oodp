package main

func main() {
	/*
		in this example we are creating a cache that has different strategys of eviction.
		Assuming the cache has a limit of memory, when attempting to insert into the cache
		and reaching the limit, the cache will utilize an algorithm to determine an item
		to evict.
	*/
	lru := strategyLRU{}
	lfu := strategyLFU{}
	fifo := strategyFIFO{}
	c := newCache(lru)
	c.add("k1", "v1")
	c.add("k2", "v2")
	c.add("k3", "v3")

	/*
		our application allows users at runtime to change the strategy
	*/
	c.setStrategy(lfu)
	c.add("k4", "v4")

	c.setStrategy(fifo)
	c.add("k5", "v5")
}
