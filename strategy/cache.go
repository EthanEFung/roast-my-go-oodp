package main

import "time"

type node struct {
	val       string
	createdAt time.Time
	updatedAt time.Time
	visited   int
}

/*
	cache is the context object of our example. It stores reference to a cache strategy
	and has a `setStrategy` method that allows the strategy to change
*/
type cache struct {
	store       map[string]*node
	strategy    strategy[*cache]
	maxCapacity int
}

func newCache(s strategy[*cache]) *cache {
	c := new(cache)
	c.store = make(map[string]*node)
	c.strategy = s
	c.maxCapacity = 2
	return c
}

func (c *cache) setStrategy(s strategy[*cache]) {
	c.strategy = s
}

func (c *cache) add(key, val string) {
	// here in the add method we check to make sure that the cache limit has not been reached
	if node, exists := c.store[key]; exists {
		node.updatedAt = time.Now()
		node.visited = 0
		return
	}
	if len(c.store) >= c.maxCapacity {
		// in case the limit is reached we evict an item from the cache
		c.evict()
	}
	// then we store the new key value pair in the cache
	c.store[key] = &node{
		val:       val,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		visited:   0,
	}
}
func (c *cache) get(key string) string {
	item := c.store[key]
	item.updatedAt = time.Now()
	item.visited += 1
	return item.val
}

func (c *cache) evict() {
	// notice that we utilize the strategy's execute method to evict
	c.strategy.execute(c)
}
